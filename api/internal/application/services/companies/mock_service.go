package companies

import (
	"context"
)

type mockService struct {
	companies GetSlimResponse
	err       error
}

func NewMockService(companies GetSlimResponse, err error) Service {
	return &mockService{
		companies,
		err,
	}
}

func (s *mockService) GetSlim(ctx context.Context) (GetSlimResponse, error) {
	return s.companies, s.err
}
