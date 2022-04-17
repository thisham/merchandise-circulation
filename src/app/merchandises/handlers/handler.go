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
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "server error", err)
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

func (merchHandler *MerchandiseHandler) GetMerchandiseByID(ectx echo.Context) error {
	merchID := ectx.Param("id")

	result, err := merchHandler.service.GetDataByID(merchID)
	if err != nil {
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "server error", err)
	}
	return utils.CreateResponse(ectx, http.StatusOK, "OK", result)
}

func (merchHandler *MerchandiseHandler) GetMerchandiseByUPC(ectx echo.Context) error {
	merchID := ectx.Param("upc")

	result, err := merchHandler.service.GetDataByID(merchID)
	if err != nil {
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "server error", err)
	}
	return utils.CreateResponse(ectx, http.StatusOK, "OK", result)
}

func (merchHandler *MerchandiseHandler) UpdateMerchandiseByID(ectx echo.Context) error {
	merchID := ectx.Param("id")
	merchData := request.UpdateMerchandiseRequest{
		ID: merchID,
	}

	if err := ectx.Bind(&merchData); err != nil {
		return utils.CreateResponse(ectx, http.StatusBadRequest, "bad request", err)
	}

	if err := merchHandler.service.UpdateDataByID(merchID, merchData.TranspileUpdateRequestToDomain()); err != nil {
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "server error", err)
	}
	return utils.CreateResponse(ectx, http.StatusNoContent, "no content")
}

func (merchHandler *MerchandiseHandler) DeleteMerchandiseByID(ectx echo.Context) error {
	merchID := ectx.Param("id")

	if err := merchHandler.service.DeleteDataByID(merchID); err != nil {
		return utils.CreateResponse(ectx, http.StatusInternalServerError, "server error", err)
	}
	return utils.CreateResponse(ectx, http.StatusNoContent, "no content")
}
