package services

import (
	"errors"
	"merchandise-circulation-api/src/app/users"
	"merchandise-circulation-api/src/utils"
	"strings"

	"github.com/google/uuid"
)

type usecase struct {
	repo users.Repositories
}

// AttemptLogin implements users.Services
func (uc *usecase) AttemptLogin(email string, password string) (token string, err error) {
	domain, err := uc.repo.SelectUserOnLogin(email)

	if err != nil {
		return "", err
	}

	if !utils.HashValidate(password, domain.Password) {
		return "", errors.New("password didn't match")
	}
	return utils.GenerateJwt(domain.ID.String())
}

// AttemptRegister implements users.Services
func (uc *usecase) AttemptRegister(user users.Domain) (token string, err error) {
	previousUser, err := uc.repo.SelectUserOnLogin(user.Email)
	var userID string

	if err != nil && !strings.Contains(err.Error(), "not found") {
		return "", err
	}

	if previousUser.ID != uuid.Nil {
		return "", errors.New("email is registered")
	}

	if userID, err = uc.repo.InsertUser(user); err != nil {
		return "", err
	}
	return utils.GenerateJwt(userID)
}

// GetAllUsers implements users.Services
func (uc *usecase) GetAllUsers() (users []users.Domain, err error) {
	return uc.repo.SelectUsers()
}

// GetUserByID implements users.Services
func (uc *usecase) GetUserByID(id string) (user users.Domain, err error) {
	return uc.repo.SelectUserByID(id)
}

func NewService(repo users.Repositories) users.Services {
	return &usecase{repo}
}
