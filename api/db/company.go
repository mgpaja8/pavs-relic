package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type Company struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (c Company) TableName() string {
	return "companies"
}

func (c Company) SelectBuilder(ctx context.Context) sq.SelectBuilder {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql := psql.Select("companies.*").From("companies")
	return sql
}

func (c Company) Select(ctx context.Context, sql sq.SelectBuilder) ([]Company, error) {
	query, args, err := sql.ToSql()
	if err != nil {
		return []Company{}, err
	}

	var companies []Company
	_, err = dbmap.Select(&companies, query, args...)

	return companies, err
}
