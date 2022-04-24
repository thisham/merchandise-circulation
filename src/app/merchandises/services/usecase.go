package services

import (
	"merchandise-circulation-api/src/app/merchandises"
)

type usecase struct {
	merchandiseData merchandises.Data
}

// GetAllMerchandises implements merchandises.Services
func (uc *usecase) GetAllMerchandises() ([]merchandises.Domain, error) {
	result, err := uc.merchandiseData.SelectAllMerchandises()
	return result, err
}

// GetMerchandiseByUPC implements merchandises.Services
func (uc *usecase) GetMerchandiseByUPC(upc string) (merchandises.Domain, error) {
	result, err := uc.merchandiseData.SelectMerchandiseByUPC(upc)
	return result, err
}

func NewService(merchandiseData merchandises.Data) merchandises.Services {
	return &usecase{
		merchandiseData: merchandiseData,
	}
}
