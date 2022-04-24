package services

import "merchandise-circulation-api/src/app/merchandise_types"

type usecase struct {
	merchandiseTypeData merchandise_types.Data
}

// GetAllMerchandiseTypes implements merchandise_types.Services
func (*usecase) GetAllMerchandiseTypes() ([]merchandise_types.Domain, error) {
	panic("unimplemented")
}

func NewService(merchandiseTypeData merchandise_types.Data) merchandise_types.Services {
	return &usecase{
		merchandiseTypeData: merchandiseTypeData,
	}
}
