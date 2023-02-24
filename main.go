package main

import (
	"fmt"
	"kelindan/pkg/database"
	"kelindan/pkg/logging"
	"kelindan/pkg/settings"
	"kelindan/pkg/utils"
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
	fmt.Println("claims: ", *claims)

}
