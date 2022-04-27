package tests

import (
	"fmt"
	merchandiseTypes "merchandise-circulation-api/src/app/merchandise_types/data"
	merchandises "merchandise-circulation-api/src/app/merchandises/data"
	"merchandise-circulation-api/src/configs"
	"merchandise-circulation-api/src/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mockMerchandiseUuid, _ = uuid.NewRandom()
	merchandiseSample      = merchandises.Merchandise{
		Model:             gorm.Model{},
		ID:                mockMerchandiseUuid,
		UPC:               "899921127777",
		Name:              "Pencil",
		Stock:             12,
		Price:             1300,
		Description:       "example.",
		MerchandiseTypeID: 1,
		MerchandiseType:   merchandiseTypes.MerchandiseType{ID: 1, Name: "ATK"},
	}
)

type SvrSuite struct {
	suite.Suite
	Echo   *echo.Echo
	Server *httptest.Server
	DB     *gorm.DB
}

func (s *SvrSuite) connectDB() *SvrSuite {
	conn := fmt.Sprintf("foreigner:admin@tcp(192.168.1.2:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mockDB)
	s.DB, _ = gorm.Open(mysql.Open(conn), &gorm.Config{})

	s.DB.AutoMigrate(merchandises.Merchandise{}, merchandiseTypes.MerchandiseType{})
	return s
}

func (s *SvrSuite) seedMerchandise() *SvrSuite {
	insert := s.DB.Create(&merchandiseSample)
	fmt.Println("errr", insert.Error)
	// require.NoError(s.T(), insert.Error)
	return s
}

func (s *SvrSuite) startServer() *SvrSuite {
	configs.LoadServerConfig(".")
	s.Server = httptest.NewServer(routes.New())
	return s
}

func (s *SvrSuite) closeServer() *SvrSuite {
	s.Server.Close()
	return s
}

func (s *SvrSuite) truncateDB() {
	s.DB.Exec("truncate table merchandises")
	s.DB.Exec("truncate table merchandise_types")
	require.NoError(s.T(), s.DB.Error)
}

func TestGetAllMerchandises(t *testing.T) {
	svr := new(SvrSuite).connectDB().seedMerchandise().startServer()
	defer svr.closeServer().truncateDB()

	t.Run("success", func(t *testing.T) {
		route := GetHttpExpect(t)
		res := route.GET("/merchandises").Expect().Status(http.StatusOK)
		res.JSON().Object().Keys().Length().Gt(0)
	})
}

// func TestGetAMerchandise(t *testing.T) {
// 	defer TearDown()()

// 	if err := database.DB.Create(&merchandiseSample).Error; err != nil {
// 		log.Fatal("create sample data error.")
// 	}

// 	route := GetHttpExpect(t)
// 	result := route.GET("/merchandises/" + merchandiseSample.UPC).Expect().Status(http.StatusOK).JSON().Object()
// 	result.Keys().Length().Gt(0)
// }
