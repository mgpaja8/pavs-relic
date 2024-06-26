package companies

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/mgpaja8/pavs-relic/internal/domain/models"
	"github.com/mgpaja8/pavs-relic/internal/infrastructure/persistance/mock"
)

func Test_NewService_ValidRepository(t *testing.T) {
	mockRepo := mock.NewCompanyRepository(map[uuid.UUID]models.Company{}, nil)

	service := NewService(mockRepo)

	assert.NotNil(t, service, "expected non-nil service")
}

func Test_NewService_NilRepository(t *testing.T) {
	assert.Panics(t, func() {
		NewService(nil)
	}, "expected panic")
}
