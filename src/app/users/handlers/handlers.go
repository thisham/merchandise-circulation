package handlers

import (
	"merchandise-circulation-api/src/app/users"
	"merchandise-circulation-api/src/app/users/handlers/request"
	"merchandise-circulation-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	service users.Services
}

func NewHandler(serv users.Services) *UserHandlers {
	return &UserHandlers{serv}
}

func (handler *UserHandlers) GetAllUsersHandler(ectx echo.Context) error {
	data, err := handler.service.GetAllUsers()

	if err != nil {
		return utils.CreateEchoErrorResponse(ectx, err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusOK, "OK", data)
}

func (handler *UserHandlers) GetUserByIDHandler(ectx echo.Context) error {
	userID := ectx.Param("id")
	data, err := handler.service.GetUserByID(userID)

	if err != nil {
		return utils.CreateEchoErrorResponse(ectx, err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusOK, "OK", data)
}

func (handler *UserHandlers) RegisterHandler(ectx echo.Context) error {
	var userData request.RegisterRequest

	if err := ectx.Bind(&userData); err != nil {
		return utils.CreateEchoResponse(ectx, http.StatusBadRequest, "Bad Request")
	}

	token, err := handler.service.AttemptRegister(userData.MapToDomain())
	if err != nil {
		return utils.CreateEchoErrorResponse(ectx, err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusCreated, "Created", token)
}

func (handler *UserHandlers) LoginHandler(ectx echo.Context) error {
	var userData request.LoginRequest

	if err := ectx.Bind(&userData); err != nil {
		return utils.CreateEchoResponse(ectx, http.StatusBadRequest, "Bad Request")
	}

	token, err := handler.service.AttemptLogin(userData.Email, userData.Password)
	if err != nil {
		if strings.Contains(err.Error(), "mismatch") {
			return utils.CreateEchoResponse(ectx, http.StatusUnauthorized,
				"Unauthorized")
		}
		return utils.CreateEchoErrorResponse(ectx, err)
	}
	return utils.CreateEchoResponse(ectx, http.StatusCreated, "Created", token)
}
