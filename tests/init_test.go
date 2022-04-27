package tests

import (
	"fmt"
	merchTypeData "merchandise-circulation-api/src/app/merchandise_types/data"
	merchData "merchandise-circulation-api/src/app/merchandises/data"
	"merchandise-circulation-api/src/routes"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ref: https://faun.pub/how-to-test-database-repository-in-golang-771b59c8084e

type TestSuite struct {
	suite.Suite
	DB               *gorm.DB
	Echo             *echo.Echo
	ConnectionString string
}

const mockDB = "merchandises"

func (s *TestSuite) SetupSuite() *TestSuite {
	// var err error

	s.ConnectionString = fmt.Sprintf("foreigner:admin@tcp(192.168.1.2:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mockDB)
	s.DB, _ = gorm.Open(mysql.Open(s.ConnectionString), &gorm.Config{})

	s.DB.AutoMigrate(merchData.Merchandise{}, merchTypeData.MerchandiseType{})
	return s
}

func (s *TestSuite) TearDownTest() {
	s.DB.Exec("truncate table merchandises; truncate table merchandise_types;")
	require.NoError(s.T(), s.Echo.Close())
}

func (s *TestSuite) TearDownSuite() {
	require.NoError(s.T(), s.DB.Migrator().DropTable("merchandises", "merchandise_types"))
}

func TestSuiteRan(t *testing.T) {
	if testing.Short() {
		t.Skip("sekip")
	}
	suite.Run(t, new(TestSuite))
}

func GetHttpExpect(t *testing.T) *httpexpect.Expect {
	return NewHttpExpect(t)
}

func NewHttpExpect(t *testing.T) *httpexpect.Expect {
	server := httptest.NewServer(routes.New())
	return httpexpect.New(t, server.URL)
}
