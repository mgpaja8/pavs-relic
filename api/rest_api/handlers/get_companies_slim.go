package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mgpaja8/pavs-relic/internal/application/services/companies"
)

func GetCompaniesSlim(service companies.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companies, err := service.GetSlim(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(companies); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
