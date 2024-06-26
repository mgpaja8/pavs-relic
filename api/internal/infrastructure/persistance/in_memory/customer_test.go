package inmemory

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
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

	repo := NewCustomerRepository(customers)

	params := repository.GetAllParams{}

	gotCustomers, err := repo.GetAll(context.Background(), params)

	assert.NoError(t, err, "expected no error")
	assert.ElementsMatch(t, []models.Customer{customer1, customer2}, gotCustomers, "expected customers to match")
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

	repo := NewCustomerRepository(customers)

	tests := []struct {
		name           string
		search         string
		expectedResult []models.Customer
	}{
		{"Case 1: Exact match 'Jane'", "Jane", []models.Customer{customer1}},
		{"Case 2: Partial match 'a'", "a", []models.Customer{customer1, customer2}},
		{"Case 3: Partial match in last name 'i'", "i", []models.Customer{customer2}},
		{"Case 4: Empty string", "", []models.Customer{customer1, customer2}},
		{"Case 5: No match 'empty'", "empty", []models.Customer{}},
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

	repo := NewCustomerRepository(customers)

	tests := []struct {
		name           string
		companyID      uuid.UUID
		expectedResult []models.Customer
	}{
		{"Case 1: Against Company A", uuid.New(), []models.Customer{}},
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
