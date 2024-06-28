package postgres

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"github.com/mgpaja8/pavs-relic/db"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

type customerRepository struct{}

func NewCustomerRepository() repository.CustomerRepository {
	return &customerRepository{}
}

func (r *customerRepository) GetAll(ctx context.Context, params repository.GetAllParams) ([]models.Customer, error) {
	var c db.Customer
	psql := c.SelectBuilder(ctx)

	if params.Search != nil && len(*params.Search) > 0 {
		search := fmt.Sprintf("%%%s%%", strings.ToLower(*params.Search))
		psql = psql.Where(sq.Or{
			sq.Like{"lower(first_name)": search},
			sq.Like{"lower(last_name)": search},
		})
	}

	if params.CompanyID != nil {
		psql = psql.Where(sq.Eq{"company_id": *params.CompanyID})
	}

	dbCustomers, err := c.Select(ctx, psql)
	if err != nil {
		return []models.Customer{}, err
	}

	return toCustomers(dbCustomers), nil
}

func toCustomer(c db.Customer) models.Customer {
	return models.NewCustomerFromDb(c.ID, c.CompanyID, valueobjects.FirstName(c.FirstName), valueobjects.LastName(c.LastName))
}

func toCustomers(dbCustomers []db.Customer) []models.Customer {
	customers := []models.Customer{}

	for _, c := range dbCustomers {
		customers = append(customers, toCustomer(c))
	}

	return customers
}
