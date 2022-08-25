package services

import "merchandise-circulation-api/src/app/users"

type usecase struct {
	repo users.Repositories
}

// AttemptLogin implements users.Services
func (uc *usecase) AttemptLogin(email string, password string) (jwt string, err error) {
	panic("unimplemented")
}

// AttemptRegister implements users.Services
func (uc *usecase) AttemptRegister(user users.Domain) (jwt string, err error) {
	panic("unimplemented")
}

// GetAllUsers implements users.Services
func (uc *usecase) GetAllUsers() (users []users.Domain, err error) {
	panic("unimplemented")
}

// GetUserByID implements users.Services
func (uc *usecase) GetUserByID(id string) (user users.Domain, err error) {
	panic("unimplemented")
}

func NewService(repo users.Repositories) users.Services {
	return &usecase{repo}
}
