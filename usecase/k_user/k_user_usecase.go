package usekuser

import (
	"context"
	"errors"
	"fmt"
	iuser "kelindan/interface/user"
	"kelindan/models"
	"time"

	"github.com/fatih/structs"
)

type useKUser struct {
	repoKuser      iuser.Repository
	contextTimeOut time.Duration
}

func NewUseKUser(a iuser.Repository, timeout time.Duration) iuser.Usecase {
	return &useKUser{repoKuser: a, contextTimeOut: timeout}
}

func (u *useKUser) GetByID(ctx context.Context, ID int) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	user, err := u.repoKuser.GetByID(ID)
	if err != nil {
		return result, err
	}

	response := map[string]interface{}{
		"user_id":   user.UserID,
		"user_name": user.UserName,
		"name":      user.Name,
		"email":     user.Email,
		"telp":      user.Telp,
		"is_active": user.IsActive,
		"user_type": user.UserType,
	}
	return response, nil
}

func (u *useKUser) Create(ctx context.Context, data *models.KUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoKuser.Create(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *useKUser) Update(ctx context.Context, ID int, data models.UpdateUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData, err := u.repoKuser.GetByAccount(data.Email, data.UserType)
	if userData.UserID != ID {
		return errors.New("email sudah terdaftar")
	}
	updateData := structs.Map(data)
	updateData["user_edit"] = data.UserName
	fmt.Printf("%v", updateData)
	err = u.repoKuser.Update(ID, updateData)
	if err != nil {
		return err
	}
	return nil
}
