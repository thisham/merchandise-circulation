package utils

import (
	"merchandise-circulation-api/src/utils/rescode"
	"net/http"

	"github.com/labstack/echo/v4"
)

type base struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func CreateEchoResponse(ectx echo.Context, httpCode int, reason string, data ...interface{}) error {
	response := base{}
	response.Meta.Status = httpCode
	response.Meta.Message = reason
	response.Data = data
	return ectx.JSON(httpCode, response)
}

func CreateEchoErrorResponse(ectx echo.Context, err error) error {
	if err.Error() == rescode.NotFound {
		return CreateEchoResponse(ectx, http.StatusNotFound, err.Error())
	}
	return CreateEchoResponse(ectx, http.StatusInternalServerError, "internal server error", err.Error())
}
