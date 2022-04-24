package data

import (
	"merchandise-circulation-api/src/app/merchandise_types"

	"gorm.io/gorm"
)

type record struct {
	DB *gorm.DB
}

// SelectAllMerchandiseTypes implements merchandise_types.Data
func (rec *record) SelectAllMerchandiseTypes() ([]merchandise_types.Domain, error) {
	var data []MerchandiseType
	var err = rec.DB.Find(&data).Error
	return mapToBatchDomain(data), err
}

func NewMySqlRecord(DB *gorm.DB) merchandise_types.Data {
	return &record{DB}
}
