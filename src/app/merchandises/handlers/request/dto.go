package request

import (
	"merchandise-circulation-api/src/app/merchandises"

	"github.com/google/uuid"
)

type NewMerchandiseRequest struct {
	UniversalProductCode string  `json:"universal_product_code"`
	Name                 string  `json:"name"`
	Stock                int     `json:"stock"`
	Price                float64 `json:"price"`
	Description          string  `json:"description"`
	MerchandiseTypeID    int     `json:"merchandise_type_id"`
}

type UpdateMerchandiseRequest struct {
	ID                   string  `json:"id"`
	UniversalProductCode string  `json:"universal_product_code"`
	Name                 string  `json:"name"`
	Stock                int     `json:"stock"`
	Price                float64 `json:"price"`
	Description          string  `json:"description"`
	MerchandiseTypeID    int     `json:"merchandise_type_id"`
}

func (request *NewMerchandiseRequest) TranspileCreateRequestToDomain() merchandises.Domain {
	return merchandises.Domain{
		UniversalProductCode: request.UniversalProductCode,
		Name:                 request.Name,
		Stock:                request.Stock,
		Price:                request.Price,
		Description:          request.Description,
		MerchandiseType: merchandises.MerchandiseTypeReference{
			ID: request.MerchandiseTypeID,
		},
	}
}

func (request *UpdateMerchandiseRequest) TranspileUpdateRequestToDomain() merchandises.Domain {
	requestId, _ := uuid.Parse(request.ID)
	return merchandises.Domain{
		Id:                   requestId,
		UniversalProductCode: request.UniversalProductCode,
		Name:                 request.Name,
		Stock:                request.Stock,
		Price:                request.Price,
		Description:          request.Description,
		MerchandiseType: merchandises.MerchandiseTypeReference{
			ID: request.MerchandiseTypeID,
		},
	}
}
