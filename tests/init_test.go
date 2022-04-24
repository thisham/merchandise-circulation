package tests

import (
	"merchandise-circulation-api/src/database"
	"merchandise-circulation-api/src/routes"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

var (
	echoHandler *echo.Echo
	server      *httptest.Server
)

func init() {
	echoHandler = routes.New()
	server = httptest.NewServer(echoHandler)
}

func TearDown() func() {
	return func() {
		database.DB.Exec("TRUNCATE TABLE merchandises.merchandises;")
	}
}

func GetHttpExpect(t *testing.T) *httpexpect.Expect {
	return NewHttpExpect(t)
}

func NewHttpExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: server.URL,
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
		Reporter: t,
	})
}
