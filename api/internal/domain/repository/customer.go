package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/models"
)

type GetAllParams struct {
	Search    *string
	CompanyID *uuid.UUID
}

type CustomerRepository interface {
	GetAll(ctx context.Context, params GetAllParams) ([]models.Customer, error)
}
