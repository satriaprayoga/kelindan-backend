package main

import (
	"kelindan/models"
	"kelindan/pkg/database"
	"kelindan/pkg/utils"
	repokuser "kelindan/repository/k_user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateKUser(t *testing.T) {
	repoKUser := repokuser.NewRepoKUser(database.Conn)

	user := &models.KUser{
		UserName: utils.RandomUserName(),
		Name:     utils.RandomName(),
		Password: utils.RandomPassword(),
		Telp:     utils.RandomPhone(),
		Email:    utils.RandomEmail(),
		IsActive: false,
		UserType: "user",
	}
	var err = repoKUser.Create(user)
	require.NoError(t, err)
}

func TestGetListKUser(t *testing.T) {
	repoKUser := repokuser.NewRepoKUser(database.Conn)
	var (
		params = models.ParamList{
			Page:       1,
			PerPage:    5,
			Search:     "",
			InitSearch: "",
			SortField:  "",
		}
	)
	userLists, err := repoKUser.GetAll(params)
	require.NoError(t, err)
	require.NotEmpty(t, userLists)

}
