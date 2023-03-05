package app

import (
	"fmt"
	"kelindan/pkg/utils"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func BindAndValid(c echo.Context, form interface{}) (int, string) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, fmt.Sprintf("invalid request params: %v", err)
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("invalid request params: %v", err)
	}
	if !check {
		return http.StatusBadRequest, MarkErrors(valid.Errors)
	}
	return http.StatusOK, "ok"
}

func MarkErrors(errors []*validation.Error) string {
	res := ""
	for _, err := range errors {
		res = fmt.Sprintf("%s %s", err.Key, err.Message)
	}

	return res
}

func GetClaims(e echo.Context) (utils.Claims, error) {
	var clm utils.Claims
	claims := e.Get("claims")
	err := mapstructure.Decode(claims, &clm)
	if err != nil {
		return clm, err
	}
	return clm, nil
}
