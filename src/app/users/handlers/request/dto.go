package request

import "merchandise-circulation-api/src/app/users"

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterRequest) MapToDomain() users.Domain {
	return users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginRequest) MapToDomain() users.Domain {
	return users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
