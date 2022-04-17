package utils

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func CreateResponse(ectx echo.Context, statusCode int, reason string, data ...interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = statusCode
	response.Meta.Message = reason
	response.Data = data
	return ectx.JSON(statusCode, response)
}
