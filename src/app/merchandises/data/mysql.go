package data

import (
	"merchandise-circulation-api/src/app/merchandises"

	"gorm.io/gorm"
)

type Record struct {
	DB *gorm.DB
}

// CheckDataExistsByUPC implements merchandises.Data
func (*Record) CheckDataExistsByUPC(UPC string) (bool, error) {
	panic("unimplemented")
}

// DeleteDataByID implements merchandises.Data
func (*Record) DeleteDataByID(id string) error {
	panic("unimplemented")
}

// SelectAllData implements merchandises.Data
func (*Record) SelectAllData() ([]merchandises.Domain, error) {
	panic("unimplemented")
}

// SelectDataByID implements merchandises.Data
func (*Record) SelectDataByID(id string) (merchandises.Domain, error) {
	panic("unimplemented")
}

// SelectDataByUPC implements merchandises.Data
func (*Record) SelectDataByUPC(UPC string) (merchandises.Domain, error) {
	panic("unimplemented")
}

// UpdateDataByID implements merchandises.Data
func (*Record) UpdateDataByID(id string, data merchandises.Domain) error {
	return nil
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
