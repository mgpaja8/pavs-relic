package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
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
