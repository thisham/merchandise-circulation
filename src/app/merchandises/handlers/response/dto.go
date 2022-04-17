package response

import (
	"merchandise-circulation-api/src/app/merchandises"
	"time"
)

type MerchandiseTypeReferenceResponse struct {
	ID   int    `json:"merchandise_type_id"`
	Name string `json:"merchandise_type_name"`
}

type MerchandiseResponse struct {
	ID                   string                           `json:"id"`
	UniversalProductCode string                           `json:"universal_product_code"`
	Name                 string                           `json:"name"`
	Stock                int                              `json:"stock"`
	Price                float64                          `json:"price"`
	Description          string                           `json:"description"`
	MerchandiseType      MerchandiseTypeReferenceResponse `json:"merchandise_type"`
	CreatedAt            time.Time                        `json:"created_at"`
	UpdatedAt            time.Time                        `json:"updated_at"`
}

func TranspileToResponse(domain merchandises.Domain) MerchandiseResponse {
	return MerchandiseResponse{
		ID:                   domain.Id.String(),
		UniversalProductCode: domain.UniversalProductCode,
		Name:                 domain.Name,
		Stock:                domain.Stock,
		Price:                domain.Price,
		Description:          domain.Description,
		MerchandiseType: MerchandiseTypeReferenceResponse{
			ID:   domain.MerchandiseType.ID,
			Name: domain.MerchandiseType.Name,
		},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func transpileToBatchOfResponse(domains []merchandises.Domain) []MerchandiseResponse {
	responses := []MerchandiseResponse{}

	for _, merch := range domains {
		responses = append(responses, TranspileToResponse(merch))
	}
	return responses
}
