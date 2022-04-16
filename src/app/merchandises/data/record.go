package data

import (
	"merchandise-circulation-api/src/app/merchandises"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Merchandise struct {
	gorm.Model
	ID                   uuid.UUID
	UniversalProductCode string
	Name                 string
	Stock                int
	Price                float64
	Description          string `gorm:"type:text"`
	MerchandiseTypeID    int
	MerchandiseType      MerchandiseTypeReference
}

type MerchandiseTypeReference struct {
	ID   int
	Name string
}

func transpileToRecord(domain merchandises.Domain) Merchandise {
	return Merchandise{
		Model: gorm.Model{
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: gorm.DeletedAt{},
		},
		ID:                   domain.Id,
		UniversalProductCode: domain.UniversalProductCode,
		Name:                 domain.Name,
		Stock:                domain.Stock,
		Price:                domain.Price,
		Description:          domain.Description,
		MerchandiseTypeID:    domain.MerchandiseType.ID,
		MerchandiseType: MerchandiseTypeReference{
			domain.MerchandiseType.ID, domain.MerchandiseType.Name,
		},
	}
}

func transpileToDomain(record Merchandise) merchandises.Domain {
	return merchandises.Domain{
		Id:                   record.ID,
		UniversalProductCode: record.UniversalProductCode,
		Name:                 record.Name,
		Stock:                record.Stock,
		Price:                record.Price,
		Description:          record.Description,
		MerchandiseType: merchandises.MerchandiseTypeReference{
			ID: record.MerchandiseType.ID, Name: record.MerchandiseType.Name,
		},
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func transpileToBatchOfDomain(record []Merchandise) []merchandises.Domain {
	domain := []merchandises.Domain{}

	for _, merchandise := range record {
		domain = append(domain, transpileToDomain(merchandise))
	}
	return domain
}
