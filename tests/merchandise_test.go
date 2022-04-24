package tests

import (
	"log"
	merchandiseTypes "merchandise-circulation-api/src/app/merchandise_types/data"
	merchandises "merchandise-circulation-api/src/app/merchandises/data"
	"merchandise-circulation-api/src/database"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	mockMerchandiseUuid, _ = uuid.NewRandom()
	merchandiseSample      = merchandises.Merchandise{
		Model:           gorm.Model{},
		ID:              mockMerchandiseUuid,
		UPC:             "899921127777",
		Name:            "Pencil",
		Stock:           12,
		Price:           1300,
		Description:     "example.",
		MerchandiseType: merchandiseTypes.MerchandiseType{ID: 1, Name: "ATK"},
	}
)

func TestGetAllMerchandises(t *testing.T) {
	defer TearDown()()

	if err := database.DB.Create(&merchandiseSample).Error; err != nil {
		log.Fatal("create sample data error.")
	}

	route := GetHttpExpect(t)
	result := route.GET("/merchandises").Expect().Status(http.StatusOK).JSON().Object()
	result.Keys().Length().Gt(0)
}

func TestGetAMerchandise(t *testing.T) {
	defer TearDown()()

	if err := database.DB.Create(&merchandiseSample).Error; err != nil {
		log.Fatal("create sample data error.")
	}

	route := GetHttpExpect(t)
	result := route.GET("/merchandises/" + merchandiseSample.UPC).Expect().Status(http.StatusOK).JSON().Object()
	result.Keys().Length().Gt(0)
}
