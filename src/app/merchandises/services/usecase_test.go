package services_test

import (
	"errors"
	"merchandise-circulation-api/src/app/merchandises"
	"merchandise-circulation-api/src/app/merchandises/mocks"
	"merchandise-circulation-api/src/app/merchandises/services"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData      mocks.Data
	entityDomain  merchandises.Domain
	entityService merchandises.Services
)

func TestMain(m *testing.M) {
	entityService = services.NewMerchandiseServices(&mockData)
	entityDomain = merchandises.Domain{
		Id:                   uuid.Must(uuid.NewRandom()),
		UniversalProductCode: "1122334455",
		Name:                 "Pensil Test",
		Stock:                12,
		Price:                12000,
		Description:          "Ini testing.",
		MerchandiseType:      merchandises.MerchandiseTypeReference{ID: 1, Name: "ATK"},
		CreatedAt:            time.Time{},
		UpdatedAt:            time.Time{},
	}
	os.Exit(m.Run())
}

func TestAddNewMerchandise(t *testing.T) {
	t.Run("valid input test", func(t *testing.T) {
		mockData.On("InsertData", mock.AnythingOfType("merchandises.Domain")).Return(entityDomain, nil).Once()
		result, err := entityService.CreateData(entityDomain)

		assert.Nil(t, err)
		assert.Contains(t, result.Name, "Pensil")
	})

	t.Run("cannot insert data test", func(t *testing.T) {
		mockData.On("InsertData", mock.AnythingOfType("merchandises.Domain")).Return(entityDomain, errors.New("error bro")).Once()
		_, err := entityService.CreateData(entityDomain)

		assert.NotNil(t, err)
	})

	t.Run("invalid data test", func(t *testing.T) {
		mockData.On("InsertData", mock.AnythingOfType("merchandises.Domain")).Return(entityDomain, errors.New("error")).Once()
		_, err := entityService.CreateData(entityDomain)

		assert.NotNil(t, err)
	})
}

func TestUpdateMerchandise(t *testing.T) {
	t.Run("valid update input", func(t *testing.T) {
		mockData.On("UpdateDataByID", mock.AnythingOfType("string"), mock.AnythingOfType("merchandises.Domain")).Return(nil).Once()
		err := entityService.UpdateDataByID(entityDomain.Id.String(), entityDomain)

		assert.Nil(t, err)
	})

	t.Run("invalid update input", func(t *testing.T) {
		entityDomain.Name = ""
		mockData.On("UpdateDataByID", mock.AnythingOfType("string"), entityDomain).Return(nil).Once()
		err := entityService.UpdateDataByID(entityDomain.Id.String(), entityDomain)

		assert.NotNil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockData.On("UpdateDataByID", mock.AnythingOfType("string"), entityDomain).Return(errors.New("error")).Once()
		err := entityService.UpdateDataByID(entityDomain.Id.String(), entityDomain)

		assert.NotNil(t, err)
	})
}

func TestDeleteMerchandise(t *testing.T) {
	t.Run("data successfully deleted", func(t *testing.T) {
		mockData.On("DeleteDataByID", entityDomain.Id.String()).Return(nil).Once()
		err := entityService.DeleteDataByID(entityDomain.Id.String())

		assert.Nil(t, err)
	})

	t.Run("data not found", func(t *testing.T) {
		mockData.On("DeleteDataByID", entityDomain.Id.String()).Return(errors.New("Data not found")).Once()
		err := entityService.DeleteDataByID(entityDomain.Id.String())

		assert.NotNil(t, err)
	})

	t.Run("database error", func(t *testing.T) {
		mockData.On("DeleteDataByID", entityDomain.Id.String()).Return(errors.New("Database Error")).Once()
		err := entityService.DeleteDataByID(entityDomain.Id.String())

		assert.NotNil(t, err)
	})
}

func TestSelectAllMerchandise(t *testing.T) {
	t.Run("got all data", func(t *testing.T) {
		mockData.On("SelectAllData").Return([]merchandises.Domain{entityDomain}, nil).Once()
		result, err := entityService.GetAllData()

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("database error", func(t *testing.T) {
		mockData.On("SelectAllData").Return([]merchandises.Domain{}, errors.New("database error")).Once()
		_, err := entityService.GetAllData()

		assert.NotNil(t, err)
	})
}

func TestSelectMerchandiseByID(t *testing.T) {
	t.Run("got requested data by id", func(t *testing.T) {
		mockData.On("SelectDataByID", entityDomain.Id.String()).Return(entityDomain, nil).Once()
		result, err := entityService.GetDataByID(entityDomain.Id.String())

		assert.Nil(t, err)
		assert.Equal(t, entityDomain.Name, result.Name)
	})

	t.Run("data not found", func(t *testing.T) {
		mockData.On("SelectDataByID", entityDomain.Id.String()).Return(merchandises.Domain{}, errors.New("Data not found")).Once()
		_, err := entityService.GetDataByID(entityDomain.Id.String())

		assert.NotNil(t, err)
	})

	t.Run("database error", func(t *testing.T) {
		mockData.On("SelectDataByID", entityDomain.Id.String()).Return(merchandises.Domain{}, errors.New("Database error")).Once()
		_, err := entityService.GetDataByID(entityDomain.Id.String())

		assert.NotNil(t, err)
	})
}

func TestSelectMerchandiseByUPC(t *testing.T) {
	t.Run("got requested data by UPC", func(t *testing.T) {
		mockData.On("SelectDataByUPC", entityDomain.UniversalProductCode).Return(entityDomain, nil).Once()
		result, err := entityService.GetDataByUPC(entityDomain.UniversalProductCode)

		assert.Nil(t, err)
		assert.Equal(t, entityDomain.Name, result.Name)
	})

	t.Run("data not found", func(t *testing.T) {
		mockData.On("SelectDataByUPC", entityDomain.UniversalProductCode).Return(merchandises.Domain{}, errors.New("Data not found")).Once()
		_, err := entityService.GetDataByUPC(entityDomain.UniversalProductCode)

		assert.NotNil(t, err)
	})

	t.Run("database error", func(t *testing.T) {
		mockData.On("SelectDataByUPC", entityDomain.UniversalProductCode).Return(merchandises.Domain{}, errors.New("Database error")).Once()
		_, err := entityService.GetDataByUPC(entityDomain.UniversalProductCode)

		assert.NotNil(t, err)
	})
}

func TestCheckDataExistsByUPC(t *testing.T) {
	t.Run("got requested data by UPC", func(t *testing.T) {
		mockData.On("CheckDataExistsByUPC", entityDomain.UniversalProductCode).Return(true, nil).Once()
		result, err := entityService.CheckDataExistsByUPC(entityDomain.UniversalProductCode)

		assert.Nil(t, err)
		assert.True(t, result)
	})

	t.Run("data not found", func(t *testing.T) {
		mockData.On("CheckDataExistsByUPC", entityDomain.UniversalProductCode).Return(false, errors.New("Data not found")).Once()
		result, err := entityService.CheckDataExistsByUPC(entityDomain.UniversalProductCode)

		assert.NotNil(t, err)
		assert.False(t, result)
	})

	t.Run("database error", func(t *testing.T) {
		mockData.On("CheckDataExistsByUPC", entityDomain.UniversalProductCode).Return(false, errors.New("Database error")).Once()
		result, err := entityService.CheckDataExistsByUPC(entityDomain.UniversalProductCode)

		assert.NotNil(t, err)
		assert.False(t, result)
	})
}
