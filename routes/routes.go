package routes

import (
	"kelindan/pkg/database"
	"kelindan/pkg/settings"
	"time"

	"github.com/labstack/echo/v4"

	_authcntrl "kelindan/controllers/auth"
	_repokuser "kelindan/repository/k_user"
	_useauth "kelindan/usecase/auth"
)

type AppRoutes struct {
	E *echo.Echo
}

func (e *AppRoutes) InitRouter() {
	timeoutContext := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second

	repoKUser := _repokuser.NewRepoKUser(database.Conn)
	useAuth := _useauth.NewUseAuth(repoKUser, timeoutContext)
	_authcntrl.NewAuthCntrl(e.E, useAuth)

}
