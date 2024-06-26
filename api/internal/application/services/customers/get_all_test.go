package customers

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
	"github.com/mgpaja8/pavs-relic/internal/infrastructure/persistance/mock"
)

func TestGetAll_Success(t *testing.T) {
	company1, _ := models.NewCompany(valueobjects.CompanyName("Company A"))
	company2, _ := models.NewCompany(valueobjects.CompanyName("Company B"))

	customer1, _ := models.NewCustomer(company1, valueobjects.FirstName("Jane"), valueobjects.LastName("Doe"))
	customer2, _ := models.NewCustomer(company2, valueobjects.FirstName("John"), valueobjects.LastName("Smith"))

	mockCustomers := map[uuid.UUID]models.Customer{
		customer1.ID(): customer1,
		customer2.ID(): customer2,
	}

	repo := mock.NewCustomerRepository(mockCustomers, nil)

	service := NewService(repo)

	// Define test cases
	tests := []struct {
		name     string
		request  GetAllRequest
		expected []models.Customer
	}{
		{
			name:     "Search by first name 'Jane'",
			request:  GetAllRequest{Search: "Jane"},
			expected: []models.Customer{customer1},
		},
		{
			name:     "Filter by company ID",
			request:  GetAllRequest{CompanyID: company1.ID().String()},
			expected: []models.Customer{customer1},
		},
		{
			name:     "No search criteria",
			request:  GetAllRequest{},
			expected: []models.Customer{customer1, customer2},
		},
		{
			name:     "Invalid company ID",
			request:  GetAllRequest{CompanyID: "invalid-uuid"},
			expected: nil, // Expected empty result due to error
		},
	}

	// Run subtests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetAll(context.Background(), tt.request)

			if tt.expected == nil {
				assert.Error(t, err, "expected error")
			} else {
				assert.NoError(t, err, "unexpected error")
				assert.ElementsMatch(t, tt.expected, resp.Customers, "expected customers to match")
			}
		})
	}
}

func Test_GetAll_ErrorFromRepository(t *testing.T) {
	expectedErr := errors.New("expected error")

	repo := mock.NewCustomerRepository(map[uuid.UUID]models.Customer{}, expectedErr)
	service := NewService(repo)

	_, err := service.GetAll(context.Background(), GetAllRequest{})

	assert.Error(t, err, "expected error")
	assert.EqualError(t, err, expectedErr.Error(), "error message mismatch")
}
