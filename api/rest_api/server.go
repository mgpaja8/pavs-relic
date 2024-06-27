package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/mgpaja8/pavs-relic/internal/application/services/companies"
	"github.com/mgpaja8/pavs-relic/internal/application/services/customers"

	inmemory "github.com/mgpaja8/pavs-relic/internal/infrastructure/persistance/in_memory"
	"github.com/mgpaja8/pavs-relic/rest_api/handlers"
)

func main() {
	router := setupRouter()

	corsHandler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":3001", corsHandler))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(jsonMiddleware)

	companiesService := getCompaniesService()
	customersService := getCustomerService()

	router.HandleFunc("/companies/slim", handlers.GetCompaniesSlim(companiesService)).Methods("GET")
	router.HandleFunc("/customers", handlers.GetCustomers(customersService)).Methods("GET")

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

func getCustomerService() customers.Service {
	customersRepo := inmemory.NewCustomerRepository(customersMap)

	service := customers.NewService(customersRepo)

	return service
}
