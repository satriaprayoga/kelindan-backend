package repokuser

import (
	"fmt"
	iuser "kelindan/interface/user"
	"kelindan/models"
	"kelindan/pkg/logging"
	"kelindan/pkg/settings"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoKUser struct {
	Conn *gorm.DB
}

func NewRepoKUser(Conn *gorm.DB) iuser.Repository {
	return &repoKUser{Conn}
}

func (db *repoKUser) GetByID(ID int) (result *models.KUser, err error) {
	var user = &models.KUser{}
	query := db.Conn.Where("user_id=?", ID).Find(user)
	log.Info(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (db *repoKUser) GetByAccount(Username string, UserType string) (result models.KUser, err error) {
	query := db.Conn.Where("(email ilike ? OR telp = ?) and user_type = ?", Username, Username, UserType).First(&result)
	log.Info(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return result, err
}

func (db *repoKUser) GetAll(queryparam models.ParamList) (result []*models.KUser, err error) {
	var (
		pageNum  = 0
		pageSize = settings.AppConfigSetting.App.PageSize
		qWhere   = ""
		orderBy  = queryparam.SortField
	)

	//pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}

	//order
	if queryparam.SortField != "" {
		orderBy = queryparam.SortField
	}

	// where
	if queryparam.InitSearch != "" {
		qWhere = queryparam.InitSearch
	}
	if queryparam.Search != "" {
		if qWhere != "" {
			qWhere += " and " + queryparam.Search
		} else {
			qWhere += queryparam.Search
		}
	}
	if pageNum >= 0 && pageSize > 0 {
		query := db.Conn.Where(qWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
		fmt.Printf("%v", query) //cath to log query string
		err = query.Error
	} else {
		query := db.Conn.Where(qWhere).Order(orderBy).Find(&result)
		fmt.Printf("%v", query) //cath to log query string
		err = query.Error
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil

}

func (db *repoKUser) Create(data *models.KUser) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
