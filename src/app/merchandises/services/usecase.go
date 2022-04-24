package services

import "merchandise-circulation-api/src/app/merchandises"

type usecase struct {
	merchandiseData merchandises.Data
}

// GetAllMerchandises implements merchandises.Services
func (merch *usecase) GetAllMerchandises() ([]merchandises.Domain, error) {
	panic("unimplemented")
}

// GetMerchandiseByUPC implements merchandises.Services
func (merch *usecase) GetMerchandiseByUPC(upc string) (merchandises.Domain, error) {
	panic("unimplemented")
}

func NewService(merchandiseData merchandises.Data) merchandises.Services {
	return &usecase{
		merchandiseData: merchandiseData,
	}
}
