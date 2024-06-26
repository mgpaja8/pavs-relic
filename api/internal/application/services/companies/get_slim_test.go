package companies

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
	"github.com/mgpaja8/pavs-relic/internal/infrastructure/persistance/mock"
	"github.com/mgpaja8/pavs-relic/pkg/maps"
)

func Test_GetSlim_Success(t *testing.T) {
	companyName1 := valueobjects.CompanyName("Company A")
	company1, _ := models.NewCompany(companyName1)

	companyName2 := valueobjects.CompanyName("Company B")
	company2, _ := models.NewCompany(companyName2)

	mockCompanies := map[uuid.UUID]models.Company{
		company1.ID(): company1,
		company2.ID(): company2,
	}

	repo := mock.NewCompanyRepository(mockCompanies, nil)
	service := NewService(repo)

	result, err := service.GetSlim(context.Background())

	assert.NoError(t, err, "expected no error")
	assert.ElementsMatch(t, maps.MapValues(mockCompanies), result.Companies, "expected companies mismatch")
}

func Test_GetSlim_ErrorFromRepository(t *testing.T) {
	expectedErr := errors.New("expected error")

	repo := mock.NewCompanyRepository(map[uuid.UUID]models.Company{}, expectedErr)
	service := NewService(repo)

	_, err := service.GetSlim(context.Background())

	assert.Error(t, err, "expected error")
	assert.EqualError(t, err, expectedErr.Error(), "error message mismatch")
}
