package response

import "merchandise-circulation-api/src/app/users"

type Response struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func MapToResponse(domain users.Domain) Response {
	return Response{
		ID:        domain.ID.String(),
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: domain.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func MapToBatchResponse(domains []users.Domain) []Response {
	var responses []Response

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
