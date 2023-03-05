package iauth

import (
	"context"
	"kelindan/models"
)

type Usecase interface {
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
}
