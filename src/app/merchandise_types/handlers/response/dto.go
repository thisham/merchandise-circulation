package response

import "merchandise-circulation-api/src/app/merchandise_types"

type MerchandiseTypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToResponse(domain *merchandise_types.Domain) MerchandiseTypeResponse {
	return MerchandiseTypeResponse{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func MapToDomain(response *MerchandiseTypeResponse) merchandise_types.Domain {
	return merchandise_types.Domain{
		ID:   response.ID,
		Name: response.Name,
	}
}
