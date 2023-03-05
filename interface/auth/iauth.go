package iauth

import (
	"context"
	"kelindan/models"
	"kelindan/pkg/utils"
)

type Usecase interface {
	Logout(ctx context.Context, Claims utils.Claims, Token string) (err error)
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
}
