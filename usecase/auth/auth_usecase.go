package useauth

import (
	"context"
	"errors"
	iauth "kelindan/interface/auth"
	iuser "kelindan/interface/user"
	"kelindan/models"
	"kelindan/pkg/redisdb"
	"kelindan/pkg/settings"
	"kelindan/pkg/utils"
	"time"
)

type useAuth struct {
	repoKUser      iuser.Repository
	contextTimeOut time.Duration
}

func NewUseAuth(a iuser.Repository, timeout time.Duration) iauth.Usecase {
	return &useAuth{repoKUser: a, contextTimeOut: timeout}
}

func (u *useAuth) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var (
		expireToken = settings.AppConfigSetting.JWTExpired
	)

	DataUser, err := u.repoKUser.GetByAccount(dataLogin.Account, dataLogin.UserType)
	if err != nil {
		return nil, errors.New("email anda belum terdaftar")
	}

	if !DataUser.IsActive {
		return nil, errors.New("account anda belum aktif. Silahkan register ulang ")
	}

	if utils.ComparePassword(DataUser.Password, utils.GetPassword(dataLogin.Password)) {
		return nil, errors.New("Password yang anda masukkan salah. Silahkan coba lagi...")
	}

	token, err := utils.GenerateToken(DataUser.UserID, dataLogin.Account, DataUser.UserType)
	if err != nil {
		return nil, err
	}

	redisdb.AddSession(token, DataUser.UserID, time.Duration(expireToken)*time.Hour)

	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.Name,
		"user_type": DataUser.UserType,
	}
	response := map[string]interface{}{
		"token":     token,
		"data_user": restUser,
		"user_type": DataUser.UserType,
	}

	return response, nil
}
