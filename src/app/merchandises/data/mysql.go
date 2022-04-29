package data

import (
	"errors"
	"merchandise-circulation-api/src/app/merchandises"
	"merchandise-circulation-api/src/utils/rescode"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type record struct {
	DB *gorm.DB
}

// CheckMerchandiseExistsByUPC implements merchandises.Data
func (rec *record) CheckMerchandiseExistsByUPC(upc string) bool {
	var foundData Merchandise
	rec.DB.Where("UPC = ?", upc).Find(&foundData)

	if nullishUuid, _ := uuid.Parse("0"); foundData.ID == nullishUuid {
		return false
	}
	return true
}

// SelectAllMerchandises implements merchandises.Data
func (rec *record) SelectAllMerchandises() ([]merchandises.Domain, error) {
	var merchs []Merchandise

	if err := rec.DB.Find(&merchs).Error; err != nil {
		return []merchandises.Domain{}, err
	}
	return mapToBatchDomain(merchs), nil
}

// SelectMerchandiseByUPC implements merchandises.Data
func (rec *record) SelectMerchandiseByUPC(upc string) (merchandises.Domain, error) {
	var foundData Merchandise
	err := rec.DB.Where("UPC = ?", upc).Find(&foundData).Error

	if err != nil {
		return merchandises.Domain{}, err
	}

	if nullishUuid, _ := uuid.Parse("0"); foundData.ID == nullishUuid {
		return merchandises.Domain{}, errors.New(rescode.NotFound)
	}

	return mapToDomain(foundData), nil
}

func NewMySqlRecord(DB *gorm.DB) merchandises.Data {
	return &record{DB}
}
