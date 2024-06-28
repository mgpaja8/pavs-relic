package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/mgpaja8/pavs-relic/db"
	"github.com/mgpaja8/pavs-relic/internal/application/services/companies"
	"github.com/mgpaja8/pavs-relic/internal/application/services/customers"

	"github.com/mgpaja8/pavs-relic/internal/infrastructure/persistance/postgres"
	"github.com/mgpaja8/pavs-relic/rest_api/handlers"
)

func main() {
	db.InitDB(db.ConnectionString())

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
	companiesRepo := postgres.NewCompanyRepository()

	service := companies.NewService(companiesRepo)

	return service
}

func getCustomerService() customers.Service {
	customersRepo := postgres.NewCustomerRepository()

	service := customers.NewService(customersRepo)

	return service
}
