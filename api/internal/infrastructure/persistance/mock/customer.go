package mock

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
)

type customerRepository struct {
	customers map[uuid.UUID]models.Customer
	err       error
}

func NewCustomerRepository(customers map[uuid.UUID]models.Customer, err error) repository.CustomerRepository {
	return &customerRepository{
		customers,
		err,
	}
}

func (r *customerRepository) GetAll(ctx context.Context, params repository.GetAllParams) ([]models.Customer, error) {
	if r.err != nil {
		return []models.Customer{}, r.err
	}

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
