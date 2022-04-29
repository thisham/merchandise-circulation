package handlers

import (
	"merchandise-circulation-api/src/app/merchandises"
	"merchandise-circulation-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MerchandiseHandler struct {
	service merchandises.Services
}

func NewHandler(serv merchandises.Services) *MerchandiseHandler {
	return &MerchandiseHandler{serv}
}

func (handler *MerchandiseHandler) GetAllMerchandisesHandler(ectx echo.Context) error {
	data, err := handler.service.GetAllMerchandises()

	if err != nil {
		return utils.CreateEchoErrorResponse(ectx, err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusOK, "OK", data)
}

func (handler *MerchandiseHandler) GetMerchandiseByUPCHandler(ectx echo.Context) error {
	merchUPC := ectx.Param("upc")
	data, err := handler.service.GetMerchandiseByUPC(merchUPC)

	if err != nil {
		return utils.CreateEchoErrorResponse(ectx, err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusOK, "OK", data)
}
