package authcntrl

import (
	"context"
	"fmt"
	iauth "kelindan/interface/auth"
	"kelindan/middleware"
	"kelindan/models"
	app "kelindan/pkg"
	"kelindan/pkg/logging"
	"kelindan/pkg/responses"
	"kelindan/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthCntrl struct {
	useAuth iauth.Usecase
}

func NewAuthCntrl(e *echo.Echo, useAuth iauth.Usecase) {
	cntrl := &AuthCntrl{
		useAuth: useAuth,
	}

	r := e.Group("/user/auth")
	r.POST("/register", cntrl.Register)

	L := e.Group("/user/auth/logout")
	L.Use(middleware.JWT)
	L.POST("", cntrl.Logout)

}

func (u *AuthCntrl) Logout(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var appE = responses.Resp{R: e}
	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	Token := e.Request().Header.Get("Authorization")
	err = u.useAuth.Logout(ctx, claims, Token)
	if err != nil {
		return appE.Error(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusOK, "Ok", nil)
}

func (u *AuthCntrl) Register(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = responses.Resp{R: e}
		form   = models.RegisterForm{}
	)

	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(utils.Stringify(form))
	if httpCode != 200 {
		return appE.Error(http.StatusBadRequest, errMsg, nil)
	}

	data, err := u.useAuth.Register(ctx, form)
	if err != nil {
		return appE.Error(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusOK, "Ok", data)
}
