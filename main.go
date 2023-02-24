package main

import (
	"fmt"
	"kelindan/pkg/database"
	"kelindan/pkg/logging"
	"kelindan/pkg/settings"
	"kelindan/pkg/utils"

	repokuser "kelindan/repository/k_user"
)

func init() {
	settings.Setup()
	database.Setup()

	logging.Setup()
}

func main() {

	fmt.Println(settings.AppConfigSetting.Server.RunMode)
	token, err := utils.GenerateToken(1, "test", "type1")
	if err != nil {
		fmt.Println("error generate token")
	} else {
		fmt.Printf("token : %v", token)
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("claims: ", *claims)

	repoUser := repokuser.NewRepoKUser(database.Conn)

	kuser, err_ := repoUser.GetByID(1)
	if err_ != nil {
		fmt.Println(err_)
	}
	fmt.Printf("%v", kuser)

}
