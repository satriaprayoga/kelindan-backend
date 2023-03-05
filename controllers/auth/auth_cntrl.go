package authcntrl

import (
	"context"
	"fmt"
	iauth "kelindan/interface/auth"
	"kelindan/middleware"
	app "kelindan/pkg"
	"kelindan/pkg/responses"
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
