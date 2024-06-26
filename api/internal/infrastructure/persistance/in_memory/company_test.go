package inmemory

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

func Test_CompanyRepository_GetAll_Success(t *testing.T) {
	companyName1 := valueobjects.CompanyName("Company A")
	company1, _ := models.NewCompany(companyName1)

	companyName2 := valueobjects.CompanyName("Company B")
	company2, _ := models.NewCompany(companyName2)

	mockCompanies := map[uuid.UUID]models.Company{
		company1.ID(): company1,
		company2.ID(): company2,
	}

	repo := NewCompanyRepository(mockCompanies)

	gotCompanies, err := repo.GetAll(context.Background())

	assert.NoError(t, err, "expected no error")

	assert.ElementsMatch(t, []models.Company{company1, company2}, gotCompanies, "expected companies to match")
}
