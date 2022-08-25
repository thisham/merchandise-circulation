package users

import (
	"time"

	"github.com/google/uuid"
)

// Domains
type Domain struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Abstractions
type Services interface {
	AttemptLogin(email, password string) (token string, err error)
	AttemptRegister(user Domain) (token string, err error)
	GetAllUsers() (users []Domain, err error)
	GetUserByID(id string) (user Domain, err error)
}

type Repositories interface {
	SelectUserOnLogin(email string) (domain Domain, err error)
	SelectUserByID(id string) (user Domain, err error)
	SelectUsers() (users []Domain, err error)
	InsertUser(user Domain) (userID string, err error)
}
