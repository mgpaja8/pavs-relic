package mock

import (
	"context"

	"github.com/google/uuid"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
	"github.com/mgpaja8/pavs-relic/pkg/maps"
)

type companyRepository struct {
	companies map[uuid.UUID]models.Company
	err       error
}

func NewCompanyRepository(companies map[uuid.UUID]models.Company, err error) repository.CompanyRepository {
	return &companyRepository{
		companies,
		err,
	}
}

func (r companyRepository) GetAll(ctx context.Context) ([]models.Company, error) {
	if r.err != nil {
		return []models.Company{}, r.err
	}

	return maps.MapValues(r.companies), nil
}
