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
	"strconv"
	"time"
)

type useAuth struct {
	repoKUser      iuser.Repository
	contextTimeOut time.Duration
}

func NewUseAuth(a iuser.Repository, timeout time.Duration) iauth.Usecase {
	return &useAuth{repoKUser: a, contextTimeOut: timeout}
}

func (u *useAuth) Logout(ctx context.Context, Claims utils.Claims, Token string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	redisdb.TruncateList(Token)

	return nil
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

func (u *useAuth) Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var User models.KUser

	CekData, err := u.repoKUser.GetByAccount(dataRegister.Account, dataRegister.UserType)

	if CekData.Email == dataRegister.Account {
		if CekData.IsActive {
			return output, errors.New("email sudah terdaftar")
		}
	}

	if dataRegister.Passwd != dataRegister.ConfirmPasswd {
		return output, errors.New("password dan confirm password harus sama")
	}

	User.Name = dataRegister.Name
	User.Password, _ = utils.Hash(dataRegister.Passwd)
	User.IsActive = false
	User.Email = dataRegister.Account
	User.UserType = dataRegister.UserType

	if CekData.UserID > 0 && !CekData.IsActive {
		CekData.Name = User.Name
		CekData.Password = User.Password
		CekData.UserType = User.UserType
		CekData.IsActive = User.IsActive
		CekData.Email = User.Email

		err = u.repoKUser.Update(CekData.UserID, CekData)
		if err != nil {
			return output, nil
		}
	} else {
		User.UserEdit = dataRegister.Name
		User.UserInput = dataRegister.Name
		err = u.repoKUser.Create(&User)
		if err != nil {
			return output, err
		}
		mUser := map[string]interface{}{
			"user_input": strconv.Itoa(User.UserID),
			"user_edit":  strconv.Itoa(User.UserID),
		}
		err = u.repoKUser.Update(User.UserID, mUser)
		if err != nil {
			return output, err
		}
	}

	GenCode := utils.GenerateNumber(4)
	if CekData.UserID > 0 {
		redisdb.TruncateList(dataRegister.Account + "_Register")
	}

	err = redisdb.AddSession(dataRegister.Account+"_Register", GenCode, 24*time.Hour)
	if err != nil {
		return output, err
	}
	out := map[string]interface{}{
		"otp":     GenCode,
		"account": User.Email,
	}
	return out, nil

}
