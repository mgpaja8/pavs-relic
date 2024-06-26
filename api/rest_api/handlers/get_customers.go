package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mgpaja8/pavs-relic/internal/application/services/customers"
)

func GetCustomers(service customers.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		companyID := r.URL.Query().Get("company_id")

		params := customers.GetAllRequest{
			Search:    search,
			CompanyID: companyID,
		}

		customers, err := service.GetAll(r.Context(), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(customers); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
