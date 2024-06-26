package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mgpaja8/pavs-relic/internal/application/services/companies"

	inmemory "github.com/mgpaja8/pavs-relic/internal/infrastructure/persistance/in_memory"
	"github.com/mgpaja8/pavs-relic/rest_api/handlers"
)

func main() {
	router := setupRouter()
	log.Fatal(http.ListenAndServe(":3001", router))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(jsonMiddleware)

	companiesService := getCompaniesService()
	router.HandleFunc("/companies/slim", handlers.GetCompaniesSlim(companiesService)).Methods("GET")

	return router
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getCompaniesService() companies.Service {
	companiesRepo := inmemory.NewCompanyRepository(companiesMap)

	service := companies.NewService(companiesRepo)

	return service
}
