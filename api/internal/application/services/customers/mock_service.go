package customers

import (
	"context"
)

type mockService struct {
	customers GetAllResponse
	err       error
}

func NewMockService(customers GetAllResponse, err error) Service {
	return &mockService{
		customers,
		err,
	}
}

func (s *mockService) GetAll(ctx context.Context, _ GetAllRequest) (GetAllResponse, error) {
	return s.customers, s.err
}
