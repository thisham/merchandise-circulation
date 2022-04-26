package data

import (
	"merchandise-circulation-api/src/app/merchandises"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type record struct {
	DB *gorm.DB
}

// CheckMerchandiseExistsByUPC implements merchandises.Data
func (rec *record) CheckMerchandiseExistsByUPC(upc string) bool {
	var foundData Merchandise
	err := rec.DB.Where("UPC = ?", upc).Find(&foundData).Error
	nullishUuid, _ := uuid.Parse("0")
	result := foundData.ID == nullishUuid || err != nil
	return result
}

// SelectAllMerchandises implements merchandises.Data
func (rec *record) SelectAllMerchandises() ([]merchandises.Domain, error) {
	var data []Merchandise

	if err := rec.DB.Find(&data).Error; err != nil {
		return []merchandises.Domain{}, nil
	}
	return mapToBatchDomain(data), nil
}

// SelectMerchandiseByUPC implements merchandises.Data
func (rec *record) SelectMerchandiseByUPC(upc string) (merchandises.Domain, error) {
	var foundData Merchandise
	err := rec.DB.Where("UPC = ?", upc).Find(&foundData).Error
	nullishUuid, _ := uuid.Parse("0")

	if foundData.ID == nullishUuid || err != nil {
		return merchandises.Domain{}, err
	}

	return mapToDomain(foundData), nil
}

func NewMySqlRecord(DB *gorm.DB) merchandises.Data {
	return &record{DB}
}
