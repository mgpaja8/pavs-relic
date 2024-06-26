package companies

import "context"

type Service interface {
	GetSlim(ctx context.Context) (GetSlimResponse, error)
}
