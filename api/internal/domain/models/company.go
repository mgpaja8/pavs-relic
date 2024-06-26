package models

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/mgpaja8/pavs-relic/internal/domain/valueobjects"
)

type Company struct {
	id   uuid.UUID
	name valueobjects.CompanyName
}

func (c Company) ID() uuid.UUID                  { return c.id }
func (c Company) Name() valueobjects.CompanyName { return c.name }

func NewCompany(name valueobjects.CompanyName) (Company, error) {
	return Company{
		id:   uuid.New(),
		name: name,
	}, nil
}

func (c Company) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   c.id.String(),
		Name: c.name.String(),
	})
}
