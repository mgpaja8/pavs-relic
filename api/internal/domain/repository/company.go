package repository

import (
	"context"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
)

type CompanyRepository interface {
	GetAll(ctx context.Context) ([]models.Company, error)
}
