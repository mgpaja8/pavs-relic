package inmemory

import (
	"context"
	"strings"
	"sync"

	"github.com/google/uuid"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
)

type customerRepository struct {
	mu        sync.RWMutex
	customers map[uuid.UUID]models.Customer
}

func NewCustomerRepository(customers map[uuid.UUID]models.Customer) repository.CustomerRepository {
	return &customerRepository{
		customers: customers,
	}
}

func (r *customerRepository) GetAll(ctx context.Context, params repository.GetAllParams) ([]models.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Customer

	for _, customer := range r.customers {
		if params.Search != nil && len(*params.Search) > 0 {
			search := strings.ToLower(*params.Search)
			firstName := strings.ToLower(customer.FirstName().String())
			lastName := strings.ToLower(customer.LastName().String())

			if !strings.Contains(firstName, search) && !strings.Contains(lastName, search) {
				continue
			}
		}

		if params.CompanyID != nil {
			if customer.CompanyID() != *params.CompanyID {
				continue
			}
		}

		result = append(result, customer)
	}

	return result, nil
}
