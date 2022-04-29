package merchandises

import (
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID              uuid.UUID
	MerchandiseType MerchandiseTypeReference
	UPC             string
	Name            string
	Stock           int
	Price           float64
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MerchandiseTypeReference struct {
	ID   int
	Name string
}

// main clauses: GET, CREATE, AMEND, REMOVE
type Services interface {
	GetAllMerchandises() ([]Domain, error)
	GetMerchandiseByUPC(upc string) (Domain, error)
}

// main clauses: INSERT, SELECT, UPDATE, DELETE
type Data interface {
	SelectAllMerchandises() ([]Domain, error)
	SelectMerchandiseByUPC(upc string) (Domain, error)
	CheckMerchandiseExistsByUPC(upc string) bool
}
