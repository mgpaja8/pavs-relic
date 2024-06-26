package inmemory

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
	"github.com/mgpaja8/pavs-relic/pkg/maps"
)

type companyRepository struct {
	mu        sync.RWMutex
	companies map[uuid.UUID]models.Company
}

func NewCompanyRepository(companies map[uuid.UUID]models.Company) repository.CompanyRepository {
	return &companyRepository{
		companies: companies,
	}
}

func (r *companyRepository) GetAll(ctx context.Context) ([]models.Company, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return maps.MapValues(r.companies), nil
}
