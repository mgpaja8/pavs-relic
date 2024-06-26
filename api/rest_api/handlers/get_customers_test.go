package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/application/services/customers"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
)

func Test_GetCustomers_Success(t *testing.T) {
	mockService := customers.NewMockService(customers.GetAllResponse{}, nil)

	req := httptest.NewRequest("GET", "/customers", nil)
	w := httptest.NewRecorder()

	handlerFunc := GetCustomers(mockService)
	handlerFunc(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "HTTP status code should be 200")

	var responseCustomers customers.GetAllResponse
	err := json.NewDecoder(w.Body).Decode(&responseCustomers)
	if err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

	expectedCustomers := []models.Customer{}
	assert.ElementsMatch(t, expectedCustomers, responseCustomers.Customers, "response should match expected customers")
}

func Test_GetCustomers_Error(t *testing.T) {
	err := errors.New("service err")

	mockService := customers.NewMockService(customers.GetAllResponse{}, err)

	req := httptest.NewRequest("GET", "/companies/slim", nil)
	w := httptest.NewRecorder()

	handlerFunc := GetCustomers(mockService)
	handlerFunc(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code, "HTTP status code should be 500")
	assert.Equal(t, "service err\n", w.Body.String(), "Error message should match mock error")
}
