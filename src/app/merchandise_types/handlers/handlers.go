package handlers

import (
	"merchandise-circulation-api/src/app/merchandise_types"
	"merchandise-circulation-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MerchandiseTypeHandler struct {
	service merchandise_types.Services
}

func NewHandler(serv merchandise_types.Services) *MerchandiseTypeHandler {
	return &MerchandiseTypeHandler{serv}
}

func (handler *MerchandiseTypeHandler) GetAllMerchandiseTypesHandler(ectx echo.Context) error {
	data, err := handler.service.GetAllMerchandiseTypes()

	if err != nil {
		return utils.CreateEchoResponse(ectx, http.StatusInternalServerError, "cannot connect db", err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusOK, "OK", data)
}
