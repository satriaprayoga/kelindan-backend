package iuser

import "kelindan/models"

type Repository interface {
	GetByID(ID int) (result *models.KUser, err error)
	//GetByUsername(username string) (result models.KUser, err error)
	//GetAll(qParams models.ParamList) (result []*models.KUser, err error)
	Create(data *models.KUser) (err error)
	//Update(ID int, data interface{}) (err error)
	//Delete(ID int) (err error)
	//UpdatePasswordByEmail(Email string, Password string) error
	//Count(qParams models.ParamList) (result int, err error)
}
