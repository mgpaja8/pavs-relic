package models

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func Test_Customer_Getters(t *testing.T) {
	customerID := uuid.New()
	companyID := uuid.New()
	firstName, _ := valueobjects.NewFirstName("John")
	lastName, _ := valueobjects.NewLastName("Doe")

	customer := Customer{
		id:        customerID,
		companyID: companyID,
		firstName: firstName,
		lastName:  lastName,
	}

	assert.Equal(t, customerID, customer.ID(), "expected ID mismatch")
	assert.Equal(t, companyID, customer.CompanyID(), "expected CompanyID mismatch")
	assert.Equal(t, firstName, customer.FirstName(), "expected FirstName mismatch")
	assert.Equal(t, lastName, customer.LastName(), "expected LastName mismatch")
}

func Test_NewCustomer(t *testing.T) {
	companyID := uuid.New()
	company := Company{id: companyID}
	firstName, _ := valueobjects.NewFirstName("John")
	lastName, _ := valueobjects.NewLastName("Doe")

	customer, err := NewCustomer(company, firstName, lastName)

	assert.NoError(t, err, "expected no error")
	assert.NotEqual(t, uuid.Nil, customer.ID(), "expected non-zero customer ID")
	assert.Equal(t, companyID, customer.CompanyID(), "expected company ID to match")
	assert.Equal(t, firstName, customer.FirstName(), "expected first name to match")
	assert.Equal(t, lastName, customer.LastName(), "expected last name to match")
}

func Test_Customer_MarshalJSON(t *testing.T) {
	id := uuid.New()
	companyID := uuid.New()
	firstName, _ := valueobjects.NewFirstName("John")
	lastName, _ := valueobjects.NewLastName("Doe")

	customer := Customer{
		id:        id,
		companyID: companyID,
		firstName: firstName,
		lastName:  lastName,
	}

	expectedJSON := `{"id":"` + id.String() + `","company_id":"` + companyID.String() + `","first_name":"John","last_name":"Doe"}`

	marshaled, err := json.Marshal(customer)
	assert.NoError(t, err, "Error should be nil")
	assert.JSONEq(t, expectedJSON, string(marshaled), "JSON should match expected")
}
