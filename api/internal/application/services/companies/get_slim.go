package companies

import (
	"context"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
)

type GetSlimResponse struct {
	Companies []models.Company
}

func (s *service) GetSlim(ctx context.Context) (GetSlimResponse, error) {
	companies, err := s.companies.GetAll(ctx)
	if err != nil {
		return GetSlimResponse{}, err
	}

	return GetSlimResponse{
		Companies: companies,
	}, nil
}
