package main

import (
	"fmt"
	"kelindan/pkg/database"
	"kelindan/pkg/logging"
	"kelindan/pkg/redisdb"
	"kelindan/pkg/settings"
	"kelindan/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	settings.Setup()
	database.Setup()
	redisdb.Setup()
	logging.Setup()
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	R := routes.AppRoutes{E: e}
	R.InitRouter()
	sPort := fmt.Sprintf(":%d", settings.AppConfigSetting.Server.HTTPPort)
	log.Fatal(e.Start(sPort))

	// timeOutCtx := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second
	// repoUser := repokuser.NewRepoKUser(database.Conn)
	// useUser := usekuser.NewUseKUser(repoUser, timeOutCtx)

	// updateUser := models.UpdateUser{
	// 	UserName: "john32",
	// 	Name:     "john tulang",
	// 	Telp:     "08131222102",
	// 	Email:    "john31@gmail.com",
	// 	UserType: "editor",
	// }

	// err_ := useUser.Update(context.Background(), 1, updateUser)
	// if err_ != nil {
	// 	fmt.Println(err_)
	// } else {
	// 	fmt.Printf("OK")
	// }

}
