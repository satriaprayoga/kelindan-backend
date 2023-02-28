package usekuser

import (
	"context"
	"errors"
	"fmt"
	iuser "kelindan/interface/user"
	"kelindan/models"
	"kelindan/pkg/database"
	"kelindan/pkg/responses"
	"kelindan/pkg/utils"
	"math"
	"reflect"
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

func (u *useKUser) GetList(ctx context.Context, queryparams models.ParamList) (result responses.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var tUser = models.KUser{}
	/*membuat Where like dari struct*/
	if queryparams.Search != "" {
		value := reflect.ValueOf(tUser)
		types := reflect.TypeOf(&tUser)
		queryparams.Search = database.GetWhereLikeStruct(value, types, queryparams.Search, "")
	}
	result.Data, err = u.repoKuser.GetAll(queryparams)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoKuser.Count(queryparams)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparams.PerPage)))
	result.Page = queryparams.Page

	return result, nil
}

func (u *useKUser) GetByEmailKUser(ctx context.Context, email string, usertype string) (result models.KUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	a := models.KUser{}
	result, err = u.repoKuser.GetByAccount(email, usertype)
	if err != nil {
		return a, err
	}
	return result, nil
}

func (u *useKUser) ChangePassword(ctx context.Context, ID int, ChangePassword models.ChangePassword) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	dataUser, err := u.repoKuser.GetByID(ID)
	if err != nil {
		return err
	}

	if !utils.ComparePassword(dataUser.Password, utils.GetPassword(ChangePassword.OldPassword)) {
		return errors.New("your old password is wrong")
	}

	if ChangePassword.NewPassword != ChangePassword.ConfirmPassword {
		return errors.New("your confirm password is wrong")
	}

	if utils.ComparePassword(dataUser.Password, utils.GetPassword(ChangePassword.NewPassword)) {
		return errors.New("new password can't be same as old password")
	}

	ChangePassword.NewPassword, _ = utils.Hash(ChangePassword.NewPassword)

	err = u.repoKuser.UpdatePasswordByEmail(dataUser.Email, ChangePassword.NewPassword, dataUser.UserType)
	if err != nil {
		return err
	}
	return nil
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

func (u *useKUser) Delete(ctx context.Context, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoKuser.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
