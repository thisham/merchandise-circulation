package response

import "merchandise-circulation-api/src/app/merchandise_types"

type MerchandiseTypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToResponse(domain merchandise_types.Domain) MerchandiseTypeResponse {
	return MerchandiseTypeResponse{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func MapToBatchResponse(domain []merchandise_types.Domain) []MerchandiseTypeResponse {
	responses := []MerchandiseTypeResponse{}

	for _, merchType := range domain {
		responses = append(responses, MerchandiseTypeResponse(merchType))
	}
	return responses
}
