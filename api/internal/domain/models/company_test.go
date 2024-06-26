package models

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

func Test_Company_Getters(t *testing.T) {
	companyID := uuid.New()
	companyName, err := valueobjects.NewCompanyName("Pavs Relic Corporation")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	company := Company{
		id:   companyID,
		name: companyName,
	}

	assert.Equal(t, companyID, company.ID(), "expected ID mismatch")
	assert.Equal(t, companyName, company.Name(), "expected Name mismatch")
}

func Test_NewCompany(t *testing.T) {
	name := valueobjects.CompanyName("Pavs Relic Company")

	company, err := NewCompany(name)

	assert.NoError(t, err, "expected no error")
	assert.NotEqual(t, uuid.Nil, company.ID(), "expected non-zero company ID")
	assert.Equal(t, name, company.Name(), "expected company name mismatch")
}

func Test_Company_MarshalJSON(t *testing.T) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	companyName := valueobjects.CompanyName("Example Company")
	company := Company{
		id:   id,
		name: companyName,
	}

	jsonBytes, err := json.Marshal(company)
	if err != nil {
		t.Errorf("Error marshaling Company to JSON: %v", err)
	}

	expectedJSON := `{"id":"123e4567-e89b-12d3-a456-426614174000","name":"Example Company"}`

	assert.JSONEq(t, expectedJSON, string(jsonBytes), "JSON output should match expected")
}
