package main

import (
	"context"
	"fmt"
	"kelindan/models"
	"kelindan/pkg/database"
	"kelindan/pkg/logging"
	"kelindan/pkg/settings"
	"time"

	repokuser "kelindan/repository/k_user"
	usekuser "kelindan/usecase/k_user"
)

func init() {
	settings.Setup()
	database.Setup()

	logging.Setup()
}

func main() {

	timeOutCtx := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second
	repoUser := repokuser.NewRepoKUser(database.Conn)
	useUser := usekuser.NewUseKUser(repoUser, timeOutCtx)

	updateUser := models.UpdateUser{
		UserName: "john32",
		Name:     "john tulang",
		Telp:     "08131222102",
		Email:    "john31@gmail.com",
		UserType: "editor",
	}

	err_ := useUser.Update(context.Background(), 1, updateUser)
	if err_ != nil {
		fmt.Println(err_)
	} else {
		fmt.Printf("OK")
	}

}
