package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/application/services/companies"
	"github.com/mgpaja8/pavs-relic/rest_api/handlers"
)

func Test_SetupRouter(t *testing.T) {
	router := setupRouter()

	mockService := companies.NewMockService(companies.GetSlimResponse{}, nil)

	router.HandleFunc("/companies/slim", handlers.GetCompaniesSlim(mockService)).Methods("GET")

	req := httptest.NewRequest("GET", "/companies/slim", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "HTTP status code should be 200")
}

func Test_JSONMiddleware(t *testing.T) {
	router := mux.NewRouter()
	router.Use(jsonMiddleware)

	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	router.HandleFunc("/dummy", dummyHandler).Methods("GET")

	req := httptest.NewRequest("GET", "/dummy", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Content-Type header should be application/json")
}
