package tests

import (
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

func TearDown() func() error {
	return func() error {
		return nil
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
