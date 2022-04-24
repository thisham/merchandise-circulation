package data

import (
	"merchandise-circulation-api/src/app/merchandises"

	"gorm.io/gorm"
)

type Record struct {
	DB *gorm.DB
}

// CheckMerchandiseExistsByUPC implements merchandises.Data
func (*Record) CheckMerchandiseExistsByUPC(upc string) bool {
	panic("unimplemented")
}

// SelectAllMerchandises implements merchandises.Data
func (*Record) SelectAllMerchandises() ([]merchandises.Domain, error) {
	panic("unimplemented")
}

// SelectMerchandiseByUPC implements merchandises.Data
func (*Record) SelectMerchandiseByUPC(upc string) (merchandises.Domain, error) {
	panic("unimplemented")
}

func NewMySqlRecord(DB *gorm.DB) merchandises.Data {
	return &Record{DB}
}
