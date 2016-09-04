package app_test

import (
	"net/http/httptest"
	"fmt"
	"calc-feed-nums/app"
	"testing"
	"net/http"
	"encoding/json"
	"calc-feed-nums/app/transport/dto"
)

var (
	server   *httptest.Server
	numbersUrl string
)

func init() {
	// Init service and its endpoints
	server = httptest.NewServer(app.Handlers())
	numbersUrl = fmt.Sprintf("%s/numbers", server.URL)
}

func TestNumbers(t *testing.T) {
	// Arrange
	resStruct := dto.Numbers{}

	// Act
	res, err := http.Get(numbersUrl)

	// Assert
	// Check error in request
	if err != nil {
		t.Error(err)
	}
	// Check json format
	err = json.NewDecoder(res.Body).Decode(&resStruct)
	if err != nil {
		t.Error(err)
	}
	// Check status code
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}