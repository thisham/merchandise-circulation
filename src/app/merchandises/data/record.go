package data

import (
	"merchandise-circulation-api/src/app/merchandise_types"
	"merchandise-circulation-api/src/app/merchandise_types/data"
	"merchandise-circulation-api/src/app/merchandises"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// model
type Merchandise struct {
	gorm.Model
	ID                uuid.UUID
	UPC               string `gorm:"unique;not null;type:varchar(32)"`
	Name              string `gorm:"unique;not null;type:varchar(100)"`
	Stock             int
	Price             float64
	Description       string `gorm:"type:text"`
	MerchandiseTypeID int    `gorm:"not null"`
	MerchandiseType   data.MerchandiseType
}

func mapToRecord(domain merchandises.Domain) Merchandise {
	return Merchandise{
		ID:                domain.ID,
		UPC:               domain.UPC,
		Name:              domain.Name,
		Stock:             domain.Stock,
		Price:             domain.Price,
		Description:       domain.Description,
		MerchandiseTypeID: domain.MerchandiseType.ID,
		MerchandiseType: data.MerchandiseType{
			ID: domain.MerchandiseType.ID, Name: domain.MerchandiseType.Name,
		},
	}
}

func mapToDomain(record Merchandise) merchandises.Domain {
	return merchandises.Domain{
		ID:          record.ID,
		UPC:         record.UPC,
		Name:        record.Name,
		Stock:       record.Stock,
		Price:       record.Price,
		Description: record.Description,
		MerchandiseType: merchandise_types.Domain{
			ID: record.MerchandiseType.ID, Name: record.MerchandiseType.Name,
		},
		CreatedAt: record.Model.CreatedAt,
		UpdatedAt: record.Model.UpdatedAt,
	}
}

func mapToBatchDomain(records []Merchandise) []merchandises.Domain {
	domain := []merchandises.Domain{}

	for _, rec := range records {
		domain = append(domain, mapToDomain(rec))
	}
	return domain
}
