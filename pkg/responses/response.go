package responses

import (
	"kelindan/pkg/logging"
	"kelindan/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	R echo.Context
}

type ResponseModel struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseModelList struct {
	Page         int         `json:"page"`
	Total        int         `json:"total"`
	LastPage     int         `json:"last_page"`
	DefineSize   string      `json:"define_size"`
	DefineColumn string      `json:"define_column"`
	AllColumn    string      `json:"all_column"`
	Data         interface{} `json:"data"`
	Msg          string      `json:"message"`
}

func (e Resp) Response(code int, errorMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{
		Msg:  errorMsg,
		Data: data,
	}

	logger.Info(string(utils.Stringify(response)))
	return e.R.JSON(code, response)
}

func (e Resp) Error(code int, errorMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{
		Msg:  errorMsg,
		Data: data,
	}

	logger.Error(string(utils.Stringify(response)))
	return e.R.JSON(code, response)
}

func (e Resp) ErrorList(code int, errMsg string, data ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Error(string(utils.Stringify(data)))
	return e.R.JSON(code, data)
}

func (e Resp) ResponseList(code int, errMsg string, data ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Info(string(utils.Stringify(data)))
	return e.R.JSON(code, data)
}
