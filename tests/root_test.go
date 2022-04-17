package tests

import (
	"net/http"
	"testing"
)

func TestGetDefaultServerStatusResponse(t *testing.T) {
	defer TearDown()
	route := GetHttpExpect(t)

	result := route.GET("/").Expect().Status(http.StatusOK)
	result.Body().Contains("runs perfectly!")
}
