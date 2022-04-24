package services

import (
	"merchandise-circulation-api/src/app/merchandise_types"
)

type usecase struct {
	merchandiseTypeData merchandise_types.Data
}

// GetAllMerchandiseTypes implements merchandise_types.Services
func (uc *usecase) GetAllMerchandiseTypes() ([]merchandise_types.Domain, error) {
	result, err := uc.merchandiseTypeData.SelectAllMerchandiseTypes()
	return result, err
}

func NewService(merchandiseTypeData merchandise_types.Data) merchandise_types.Services {
	return &usecase{
		merchandiseTypeData: merchandiseTypeData,
	}
}
