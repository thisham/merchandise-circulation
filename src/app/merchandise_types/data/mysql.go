package data

import (
	"merchandise-circulation-api/src/app/merchandise_types"

	"gorm.io/gorm"
)

type record struct {
	DB *gorm.DB
}

// SelectAllMerchandiseTypes implements merchandise_types.Data
func (*record) SelectAllMerchandiseTypes() (merchandise_types.Domain, error) {
	panic("unimplemented")
}

func NewMySqlRecord(DB *gorm.DB) merchandise_types.Data {
	return &record{DB}
}
