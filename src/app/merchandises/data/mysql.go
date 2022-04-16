package data

import (
	"merchandise-circulation-api/src/app/merchandises"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	DB *gorm.DB
}

// CheckDataExistsByUPC implements merchandises.Data
func (rec *Record) CheckDataExistsByUPC(UPC string) (bool, error) {
	var foundData Merchandise
	err := rec.DB.Where("universal_product_code = ?", UPC).Find(&foundData).Error
	nullishUuid, _ := uuid.Parse("0")

	if foundData.ID == nullishUuid || err != nil {
		return false, err
	}

	return true, nil
}

// DeleteDataByID implements merchandises.Data
func (rec *Record) DeleteDataByID(id string) error {
	var merchModel Merchandise
	err := rec.DB.Where("ID = ?", id).Delete(&merchModel).Error
	return err
}

// SelectAllData implements merchandises.Data
func (rec *Record) SelectAllData() ([]merchandises.Domain, error) {
	var allData []Merchandise
	err := rec.DB.Find(&allData).Error

	if err != nil {
		return nil, err
	}

	return transpileToBatchOfDomain(allData), nil
}

// SelectDataByID implements merchandises.Data
func (rec *Record) SelectDataByID(id string) (merchandises.Domain, error) {
	var foundData Merchandise
	err := rec.DB.Where("ID = ?", id).Find(&foundData).Error
	nullishUuid, _ := uuid.Parse("0")

	if foundData.ID == nullishUuid || err != nil {
		return merchandises.Domain{}, err
	}

	return transpileToDomain(foundData), nil
}

// SelectDataByUPC implements merchandises.Data
func (rec *Record) SelectDataByUPC(UPC string) (merchandises.Domain, error) {
	var foundData Merchandise
	err := rec.DB.Where("universal_product_code = ?", UPC).Find(&foundData).Error
	nullishUuid, _ := uuid.Parse("0")

	if foundData.ID == nullishUuid || err != nil {
		return merchandises.Domain{}, err
	}

	return transpileToDomain(foundData), nil
}

// UpdateDataByID implements merchandises.Data
func (rec *Record) UpdateDataByID(id string, data merchandises.Domain) error {
	var merchModel Merchandise
	convertedData := transpileToRecord(data)
	err := rec.DB.Model(&merchModel).Where("ID = ?", id).Updates(&convertedData).Error
	return err
}

// InsertData implements merchandises.Data
func (rec *Record) InsertData(data merchandises.Domain) (merchandises.Domain, error) {
	convertedData := transpileToRecord(data)

	if err := rec.DB.Create(&convertedData).Error; err != nil {
		return merchandises.Domain{}, err
	}

	return transpileToDomain(convertedData), nil
}

func NewMySqlRecord(DB *gorm.DB) merchandises.Data {
	return &Record{DB}
}
