package iuser

import (
	"context"
	"kelindan/models"
	"kelindan/pkg/utils"
)

type Repository interface {
	GetByID(ID int) (result *models.KUser, err error)
	GetByAccount(Username string, UserType string) (result models.KUser, err error)
	GetAll(queryparams models.ParamList) (result []*models.KUser, err error)
	Create(data *models.KUser) (err error)
	Update(ID int, data interface{}) (err error)
	Delete(ID int) (err error)
	UpdatePasswordByEmail(Email string, Password string) error
	Count(queryparams models.ParamList) (result int, err error)
}

type Usecase interface {
	GetByID(ctx context.Context, Claims utils.Claims, ID int) (result interface{}, err error)
}
