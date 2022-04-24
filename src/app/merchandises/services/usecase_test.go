package services_test

import (
	"errors"
	"merchandise-circulation-api/src/app/merchandise_types"
	"merchandise-circulation-api/src/app/merchandises"
	"merchandise-circulation-api/src/app/merchandises/mocks"
	"merchandise-circulation-api/src/app/merchandises/services"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockData      mocks.Data
	entityDomain  merchandises.Domain
	entityService merchandises.Services
)

func TestMain(m *testing.M) {
	entityService = services.NewService(&mockData)
	entityDomain = merchandises.Domain{
		ID:          uuid.Must(uuid.NewRandom()),
		UPC:         "112233445566",
		Name:        "Pencil Test",
		Price:       12000,
		Stock:       12,
		Description: "Testing merchandise.",
		MerchandiseType: merchandise_types.Domain{
			ID: 1, Name: "ATK",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	os.Exit(m.Run())
}

func TestGetAllMerchandises(t *testing.T) {
	t.Run("success got all data", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return([]merchandises.Domain{entityDomain}, nil).Once()
		result, err := entityService.GetAllMerchandises()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("can not connect to database", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return([]merchandises.Domain{}, errors.New("cannot connect db")).Once()
		_, err := entityService.GetAllMerchandises()

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return([]merchandises.Domain{}, nil).Once()
		result, err := entityService.GetAllMerchandises()

		assert.Nil(t, err)
		assert.Equal(t, len(result), 0)
	})
}

func TestGetAMerchandiseByUPC(t *testing.T) {
	t.Run("success got selected data", func(t *testing.T) {
		mockData.On("SelectMerchandiseByUPC", entityDomain.UPC).Return(entityDomain, nil).Once()
		result, err := entityService.GetMerchandiseByUPC(entityDomain.UPC)

		assert.Nil(t, err)
		assert.Equal(t, result.Name, entityDomain.Name)
	})

	t.Run("can not connect to database", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return(merchandises.Domain{}, errors.New("cannot connect db")).Once()
		_, err := entityService.GetMerchandiseByUPC(entityDomain.UPC)

		assert.NotNil(t, err)
	})

	t.Run("data not found", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return(merchandises.Domain{}, errors.New("data not found")).Once()
		_, err := entityService.GetMerchandiseByUPC(entityDomain.UPC)

		assert.NotNil(t, err)
	})
}
