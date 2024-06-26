package customers

import (
	"context"

	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
)

type GetAllRequest struct {
	Search    string
	CompanyID string
}

type GetAllResponse struct {
	Customers []models.Customer `json:"customers"`
}

func (s *service) GetAll(ctx context.Context, req GetAllRequest) (GetAllResponse, error) {
	filters, err := computeFilters(req)
	if err != nil {
		return GetAllResponse{}, err
	}

	customers, err := s.customers.GetAll(ctx, filters)
	if err != nil {
		return GetAllResponse{}, err
	}

	return GetAllResponse{
		Customers: customers,
	}, nil
}

func computeFilters(req GetAllRequest) (repository.GetAllParams, error) {
	filters := repository.GetAllParams{}

	if len(req.Search) > 0 {
		filters.Search = &req.Search
	}

	if len(req.CompanyID) > 0 {
		companyID, err := uuid.Parse(req.CompanyID)
		if err != nil {
			return repository.GetAllParams{}, err
		}

		filters.CompanyID = &companyID
	}

	return filters, nil
}
