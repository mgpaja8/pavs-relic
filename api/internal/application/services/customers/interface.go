package customers

import "context"

type Service interface {
	GetAll(ctx context.Context, req GetAllRequest) (GetAllResponse, error)
}
