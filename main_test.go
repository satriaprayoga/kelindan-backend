package main

import (
	"kelindan/models"
	"kelindan/pkg/database"
	"kelindan/pkg/utils"
	repokuser "kelindan/repository/k_user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
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
