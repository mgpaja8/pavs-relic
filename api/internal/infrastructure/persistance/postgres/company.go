package postgres

import (
	"context"

	"github.com/mgpaja8/pavs-relic/db"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

type companyRepository struct{}

func NewCompanyRepository() repository.CompanyRepository {
	return &companyRepository{}
}

func (r *companyRepository) GetAll(ctx context.Context) ([]models.Company, error) {
	var c db.Company
	psql := c.SelectBuilder(ctx)

	dbCompanies, err := c.Select(ctx, psql)
	if err != nil {
		return []models.Company{}, err
	}

	return toCompanies(dbCompanies), nil
}

func toCompany(c db.Company) models.Company {
	return models.NewCompanyFromDb(c.ID, valueobjects.CompanyName(c.Name))
}

func toCompanies(dbCompanies []db.Company) []models.Company {
	companies := []models.Company{}

	for _, c := range dbCompanies {
		companies = append(companies, toCompany(c))
	}

	return companies
}
