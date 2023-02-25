package usekuser

import (
	"context"
	iuser "kelindan/interface/user"
	"kelindan/pkg/utils"
	"time"
)

type useKUser struct {
	repoKuser      iuser.Repository
	contextTimeOut time.Duration
}

func NewUseKUser(a iuser.Repository, timeout time.Duration) iuser.Usecase {
	return &useKUser{repoKuser: a, contextTimeOut: timeout}
}

func (u *useKUser) GetByID(ctx context.Context, Claims utils.Claims, ID int) (result interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	user, err := u.repoKuser.GetByID(ID)
	if err != nil {
		return result, err
	}

	response := map[string]interface{}{
		"user_id": user.UserID,
	}
	return response, nil
}
