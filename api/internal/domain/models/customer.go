package models

import (
	"github.com/google/uuid"
	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

type Customer struct {
	id        uuid.UUID
	companyID uuid.UUID
	firstName valueobjects.FirstName
	lastName  valueobjects.LastName
}

func (c Customer) ID() uuid.UUID                     { return c.id }
func (c Customer) CompanyID() uuid.UUID              { return c.companyID }
func (c Customer) FirstName() valueobjects.FirstName { return c.firstName }
func (c Customer) LastName() valueobjects.LastName   { return c.lastName }

func NewCustomer(company Company, firstName valueobjects.FirstName, lastName valueobjects.LastName) (Customer, error) {
	return Customer{
		id:        uuid.New(),
		companyID: company.id,
		firstName: firstName,
		lastName:  lastName,
	}, nil
}
