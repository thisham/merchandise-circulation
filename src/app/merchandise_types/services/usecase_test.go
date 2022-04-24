package services_test

import (
	"errors"
	"merchandise-circulation-api/src/app/merchandise_types"
	"merchandise-circulation-api/src/app/merchandise_types/mocks"
	"merchandise-circulation-api/src/app/merchandise_types/services"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockData      mocks.Data
	entityDomain  merchandise_types.Domain
	entityService merchandise_types.Services
)

func TestMain(m *testing.M) {
	entityService = services.NewService(&mockData)
	entityDomain = merchandise_types.Domain{
		ID:   1,
		Name: "ATK",
	}
	os.Exit(m.Run())
}

func TestGetAllMerchandiseTypes(t *testing.T) {
	t.Run("success all got data", func(t *testing.T) {
		mockData.On("SelectAllMerchandiseTypes").Return(entityDomain, nil).Once()
		result, err := entityService.GetAllMerchandiseTypes()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("can not connect to database", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return([]merchandise_types.Domain{}, errors.New("cannot connect db")).Once()
		_, err := entityService.GetAllMerchandiseTypes()

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockData.On("SelectAllMerchandises").Return([]merchandise_types.Domain{}, nil).Once()
		result, err := entityService.GetAllMerchandiseTypes()

		assert.Nil(t, err)
		assert.Equal(t, len(result), 0)
	})
}
