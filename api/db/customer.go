package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `db:"id"`
	CompanyID uuid.UUID `db:"company_id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
}

func (c Customer) TableName() string {
	return "customers"
}

func (c Customer) SelectBuilder(ctx context.Context) sq.SelectBuilder {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql := psql.Select("customers.*").From("customers")
	return sql
}

func (c Customer) Select(ctx context.Context, sql sq.SelectBuilder) ([]Customer, error) {
	query, args, err := sql.ToSql()
	if err != nil {
		return []Customer{}, err
	}

	var customers []Customer
	_, err = dbmap.Select(&customers, query, args...)

	return customers, err
}
