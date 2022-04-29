package response

import (
	"merchandise-circulation-api/src/app/merchandise_types/handlers/response"
	"merchandise-circulation-api/src/app/merchandises"
	"time"
)

type MerchandiseResponse struct {
	ID              string                           `json:"id"`
	UPC             string                           `json:"upc"`
	Name            string                           `json:"name"`
	Stock           int                              `json:"stock"`
	Price           float64                          `json:"price"`
	Description     string                           `json:"description"`
	CreatedAt       time.Time                        `json:"created_at"`
	UpdatedAt       time.Time                        `json:"updated_at"`
	MerchandiseType response.MerchandiseTypeResponse `json:"merchandise_type"`
}

func MapToResponse(domain merchandises.Domain) MerchandiseResponse {
	return MerchandiseResponse{
		ID:          domain.ID.String(),
		UPC:         domain.UPC,
		Name:        domain.Name,
		Stock:       domain.Stock,
		Price:       domain.Price,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		MerchandiseType: response.MerchandiseTypeResponse{
			ID:   domain.MerchandiseType.ID,
			Name: domain.MerchandiseType.Name,
		},
	}
}

func MapToBatchResponse(domain []merchandises.Domain) []MerchandiseResponse {
	responses := []MerchandiseResponse{}

	for _, merch := range domain {
		responses = append(responses, MapToResponse(merch))
	}
	return responses
}
