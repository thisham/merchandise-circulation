package services

import "merchandise-circulation-api/src/app/merchandises"

type merchandisesUsecase struct {
	merchandiseData merchandises.Data
}

// CheckDataExistsByUPC implements merchandises.Services
func (merch *merchandisesUsecase) CheckDataExistsByUPC(UPC string) (bool, error) {
	result, err := merch.merchandiseData.CheckDataExistsByUPC(UPC)
	return result, err
}

// DeleteDataByID implements merchandises.Services
func (merch *merchandisesUsecase) DeleteDataByID(id string) error {
	err := merch.merchandiseData.DeleteDataByID(id)
	return err
}

// GetAllData implements merchandises.Services
func (merch *merchandisesUsecase) GetAllData() ([]merchandises.Domain, error) {
	result, err := merch.merchandiseData.SelectAllData()
	return result, err
}

// GetDataByID implements merchandises.Services
func (merch *merchandisesUsecase) GetDataByID(id string) (merchandises.Domain, error) {
	result, err := merch.merchandiseData.SelectDataByID(id)
	if err != nil {
		return merchandises.Domain{}, err
	}
	return result, nil
}

// GetDataByUPC implements merchandises.Services
func (merch *merchandisesUsecase) GetDataByUPC(UPC string) (merchandises.Domain, error) {
	result, err := merch.merchandiseData.SelectDataByUPC(UPC)
	if err != nil {
		return merchandises.Domain{}, err
	}
	return result, nil
}

// UpdateDataByID implements merchandises.Services
func (merch *merchandisesUsecase) UpdateDataByID(id string, data merchandises.Domain) error {
	response := merch.merchandiseData.UpdateDataByID(id, data)
	return response
}

// CreateData implements merchandises.Services
func (merch *merchandisesUsecase) CreateData(data merchandises.Domain) (merchandises.Domain, error) {
	result, err := merch.merchandiseData.InsertData(data)
	if err != nil {
		return merchandises.Domain{}, err
	}
	return result, err
}

func NewMerchandiseServices(merchandiseData merchandises.Data) merchandises.Services {
	return &merchandisesUsecase{
		merchandiseData: merchandiseData,
	}
}
