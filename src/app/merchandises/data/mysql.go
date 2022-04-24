package data

import (
	"merchandise-circulation-api/src/app/merchandises"

	"gorm.io/gorm"
)

type record struct {
	DB *gorm.DB
}

// CheckMerchandiseExistsByUPC implements merchandises.Data
func (*record) CheckMerchandiseExistsByUPC(upc string) bool {
	panic("unimplemented")
}

// SelectAllMerchandises implements merchandises.Data
func (*record) SelectAllMerchandises() ([]merchandises.Domain, error) {
	panic("unimplemented")
}

// SelectMerchandiseByUPC implements merchandises.Data
func (*record) SelectMerchandiseByUPC(upc string) (merchandises.Domain, error) {
	panic("unimplemented")
}

func NewMySqlRecord(DB *gorm.DB) merchandises.Data {
	return &record{DB}
}
