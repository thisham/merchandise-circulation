package merchandises

import (
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	Id                   uuid.UUID
	UniversalProductCode string
	Name                 string
	Stock                int
	Price                float64
	Description          string
	MerchandiseType      MerchandiseTypeReference
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type MerchandiseTypeReference struct {
	ID   int
	Name string
}

type Services interface {
	CreateData(data Domain) (Domain, error)
	GetAllData() ([]Domain, error)
	GetDataByID(id string) (Domain, error)
	GetDataByUPC(UPC string) (Domain, error)
	CheckDataExistsByUPC(UPC string) (bool, error)
	UpdateDataByID(id string, data Domain) error
	DeleteDataByID(id string) error
}

type Data interface {
	InsertData(data Domain) (Domain, error)
	SelectAllData() ([]Domain, error)
	SelectDataByID(id string) (Domain, error)
	SelectDataByUPC(UPC string) (Domain, error)
	CheckDataExistsByUPC(UPC string) (bool, error)
	UpdateDataByID(id string, data Domain) error
	DeleteDataByID(id string) error
}
