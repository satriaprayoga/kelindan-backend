package repokuser

import (
	"fmt"
	iuser "kelindan/interface/user"
	"kelindan/models"
	"kelindan/pkg/logging"

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
