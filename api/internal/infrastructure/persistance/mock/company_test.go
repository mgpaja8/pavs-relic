package mock

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
	"github.com/mgpaja8/pavs-relic/pkg/maps"
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

	repo := NewCompanyRepository(mockCompanies, nil)

	gotCompanies, err := repo.GetAll(context.Background())

	assert.NoError(t, err, "expected no error")

	assert.ElementsMatch(t, maps.MapValues(mockCompanies), gotCompanies, "expected companies to match")
}

func Test_CompanyRepository_GetAll_Error(t *testing.T) {
	expectedErr := errors.New("repository error")

	repo := NewCompanyRepository(map[uuid.UUID]models.Company{}, expectedErr)

	gotCompanies, err := repo.GetAll(context.Background())

	assert.Error(t, err, "expected error")
	assert.Empty(t, gotCompanies, "expected empty companies on error")
}
