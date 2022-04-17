package handlers

import (
	"merchandise-circulation-api/src/app/merchandises"
	"merchandise-circulation-api/src/app/merchandises/handlers/request"
	"merchandise-circulation-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MerchandiseHandler struct {
	service merchandises.Services
}

func NewMerchandiseHandler(serv merchandises.Services) *MerchandiseHandler {
	return &MerchandiseHandler{serv}
}

func (merchHandler *MerchandiseHandler) CreateMerchandiseHandler(ectx echo.Context) error {
	newMerch := request.NewMerchandiseRequest{}

	if err := ectx.Bind(&newMerch); err != nil {
		return utils.CreateResponse(ectx, http.StatusBadRequest, "bad request", err)
	}

	_, err := merchHandler.service.CreateData(newMerch.TranspileCreateRequestToDomain())
	if err != nil {
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "bad request", err)
	}

	return utils.CreateResponse(ectx, http.StatusCreated, "data created", newMerch)
}

func (merchHandler *MerchandiseHandler) GetAllMerchandiseHandler(ectx echo.Context) error {
	data, err := merchHandler.service.GetAllData()

	if err != nil {
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "server error", err)
	}
	return utils.CreateResponse(ectx, http.StatusOK, "OK", data)
}
