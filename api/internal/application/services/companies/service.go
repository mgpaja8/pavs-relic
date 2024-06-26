package companies

import (
	"github.com/mgpaja8/pavs-relic/internal/domain/repository"
)

type service struct {
	companies repository.CompanyRepository
}

func NewService(companies repository.CompanyRepository) Service {
	if companies == nil {
		panic("companies repo required")
	}

	return &service{
		companies,
	}
}
