package merchandises

import (
	"merchandise-circulation-api/src/app/merchandise_types"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID              uuid.UUID
	MerchandiseType merchandise_types.Domain
	UPC             string
	Name            string
	Stock           int
	Price           float64
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
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
