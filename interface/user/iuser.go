package iuser

import (
	"context"
	"kelindan/models"
	"kelindan/pkg/responses"
)

type Repository interface {
	GetByID(ID int) (result *models.KUser, err error)
	GetByAccount(Username string, UserType string) (result models.KUser, err error)
	GetAll(queryparams models.ParamList) (result []*models.KUser, err error)
	Create(data *models.KUser) (err error)
	Update(ID int, data interface{}) (err error)
	Delete(ID int) (err error)
	UpdatePasswordByEmail(Email string, Password string, UserType string) error
	Count(queryparams models.ParamList) (result int, err error)
}

type Usecase interface {
	GetByEmailKUser(ctx context.Context, email string, usertype string) (result models.KUser, err error)
	GetByID(ctx context.Context, ID int) (result interface{}, err error)
	Create(ctx context.Context, data *models.KUser) error
	Update(ctx context.Context, ID int, data models.UpdateUser) (err error)
	Delete(ctx context.Context, ID int) (err error)
	ChangePassword(ctx context.Context, ID int, ChangePassword models.ChangePassword) (err error)
	GetList(ctx context.Context, queryparams models.ParamList) (result responses.ResponseModelList, err error)
}
