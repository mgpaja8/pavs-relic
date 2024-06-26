package customers

import "github.com/mgpaja8/pavs-relic/internal/domain/repository"

type service struct {
	customers repository.CustomerRepository
}

func NewService(customers repository.CustomerRepository) Service {
	if customers == nil {
		panic("customers repo required")
	}

	return &service{
		customers,
	}
}
