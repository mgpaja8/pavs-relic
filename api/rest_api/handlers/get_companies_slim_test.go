package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/application/services/companies"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
)

func Test_GetCompaniesSlim_Success(t *testing.T) {
	mockService := companies.NewMockService(companies.GetSlimResponse{}, nil)

	req := httptest.NewRequest("GET", "/companies/slim", nil)
	w := httptest.NewRecorder()

	handlerFunc := GetCompaniesSlim(mockService)
	handlerFunc(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "HTTP status code should be 200")

	var responseCompanies companies.GetSlimResponse
	err := json.NewDecoder(w.Body).Decode(&responseCompanies)
	if err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

	expectedCompanies := []models.Company{}
	assert.ElementsMatch(t, expectedCompanies, responseCompanies.Companies, "Response should match expected companies")
}

func Test_GetCompaniesSlim_Error(t *testing.T) {
	err := errors.New("service err")

	mockService := companies.NewMockService(companies.GetSlimResponse{}, err)

	req := httptest.NewRequest("GET", "/companies/slim", nil)
	w := httptest.NewRecorder()

	handlerFunc := GetCompaniesSlim(mockService)
	handlerFunc(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code, "HTTP status code should be 500")
	assert.Equal(t, "service err\n", w.Body.String(), "Error message should match mock error")
}
