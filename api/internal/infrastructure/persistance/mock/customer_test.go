package mock

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
	"github.com/mgpaja8/pavs-relic/pkg/maps"
	"github.com/stretchr/testify/assert"
)

func Test_CustomerRepository_GetAll_Success(t *testing.T) {
	companyName1 := valueobjects.CompanyName("Company A")
	company1, _ := models.NewCompany(companyName1)

	companyName2 := valueobjects.CompanyName("Company B")
	company2, _ := models.NewCompany(companyName2)

	firstName1 := valueobjects.FirstName("Jane")
	lastName1 := valueobjects.LastName("Doe")
	customer1, _ := models.NewCustomer(company1, firstName1, lastName1)

	firstName2 := valueobjects.FirstName("Pav")
	lastName2 := valueobjects.LastName("Milicevic")
	customer2, _ := models.NewCustomer(company2, firstName2, lastName2)

	customers := map[uuid.UUID]models.Customer{
		customer1.ID(): customer1,
		customer2.ID(): customer2,
	}

	repo := NewCustomerRepository(customers, nil)

	params := repository.GetAllParams{}

	gotCustomers, err := repo.GetAll(context.Background(), params)

	assert.NoError(t, err, "expected no error")
	assert.ElementsMatch(t, maps.MapValues(customers), gotCustomers, "expected customers to match")
}

func Test_CustomerRepository_GetAll_Error(t *testing.T) {
	expectedErr := errors.New("repository error")

	repo := NewCustomerRepository(map[uuid.UUID]models.Customer{}, expectedErr)

	params := repository.GetAllParams{}

	gotCustomers, err := repo.GetAll(context.Background(), params)

	assert.Error(t, err, "expected error")
	assert.Empty(t, gotCustomers, "expected empty customers on error")
}

func Test_CustomerRepository_GetAll_WithSearch_Success(t *testing.T) {
	companyName1 := valueobjects.CompanyName("Company A")
	company1, _ := models.NewCompany(companyName1)

	companyName2 := valueobjects.CompanyName("Company B")
	company2, _ := models.NewCompany(companyName2)

	firstName1 := valueobjects.FirstName("Jane")
	lastName1 := valueobjects.LastName("Doe")
	customer1, _ := models.NewCustomer(company1, firstName1, lastName1)

	firstName2 := valueobjects.FirstName("Pav")
	lastName2 := valueobjects.LastName("Milicevic")
	customer2, _ := models.NewCustomer(company2, firstName2, lastName2)

	customers := map[uuid.UUID]models.Customer{
		customer1.ID(): customer1,
		customer2.ID(): customer2,
	}

	repo := NewCustomerRepository(customers, nil)

	tests := []struct {
		name           string
		search         string
		expectedResult []models.Customer
	}{
		{"Case 1: Exact match 'Jane'", "Jane", []models.Customer{customer1}},
		{"Case 2: Partial match 'a'", "a", []models.Customer{customer1, customer2}},
		{"Case 3: Empty string", "", []models.Customer{customer1, customer2}},
		{"Case 4: No match 'empty'", "empty", []models.Customer{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := repository.GetAllParams{
				Search: &tt.search,
			}

			gotCustomers, err := repo.GetAll(context.Background(), params)

			assert.NoError(t, err, "expected no error")
			assert.ElementsMatch(t, tt.expectedResult, gotCustomers, "expected customers to match")
		})
	}
}

func Test_CustomerRepository_GetAll_WithCompanyID_Success(t *testing.T) {
	companyName1 := valueobjects.CompanyName("Company A")
	company1, _ := models.NewCompany(companyName1)

	companyName2 := valueobjects.CompanyName("Company B")
	company2, _ := models.NewCompany(companyName2)

	firstName1 := valueobjects.FirstName("Jane")
	lastName1 := valueobjects.LastName("Doe")
	customer1, _ := models.NewCustomer(company1, firstName1, lastName1)

	firstName2 := valueobjects.FirstName("Pav")
	lastName2 := valueobjects.LastName("Milicevic")
	customer2, _ := models.NewCustomer(company2, firstName2, lastName2)

	customers := map[uuid.UUID]models.Customer{
		customer1.ID(): customer1,
		customer2.ID(): customer2,
	}

	repo := NewCustomerRepository(customers, nil)

	tests := []struct {
		name           string
		companyID      uuid.UUID
		expectedResult []models.Customer
	}{
		{"Case 1: Against Company A", company1.ID(), []models.Customer{customer1}},
		{"Case 2: Against Company B", company2.ID(), []models.Customer{customer2}},
		{"Case 3: Against Company C", uuid.New(), []models.Customer{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := repository.GetAllParams{
				CompanyID: &tt.companyID,
			}

			gotCustomers, err := repo.GetAll(context.Background(), params)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			assert.ElementsMatch(t, tt.expectedResult, gotCustomers, "expected customers to match")
		})
	}
}
