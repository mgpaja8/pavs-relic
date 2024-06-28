package models

import (
	"encoding/json"

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

func NewCustomerFromDb(id, companyID uuid.UUID, firstName valueobjects.FirstName, lastName valueobjects.LastName) Customer {
	return Customer{
		id:        id,
		companyID: companyID,
		firstName: firstName,
		lastName:  lastName,
	}
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        string `json:"id"`
		CompanyID string `json:"company_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{
		ID:        c.id.String(),
		CompanyID: c.companyID.String(),
		FirstName: c.firstName.String(),
		LastName:  c.lastName.String(),
	})
}
