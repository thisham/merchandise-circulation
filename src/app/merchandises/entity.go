package merchandises

import (
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	Id              uuid.UUID
	ProductionCode  string
	Name            string
	Stock           int
	Price           float64
	Description     string
	MerchandiseType MerchandiseTypeReference
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MerchandiseTypeReference struct {
	Name string
}

type Services interface {
	CreateData(data Domain) (Domain, error)
	GetAllData() []Domain
}

type Data interface {
	InsertData(data Domain) (Domain, error)
	SelectData() []Domain
}
