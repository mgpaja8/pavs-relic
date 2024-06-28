package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

var dbmap *gorp.DbMap
var dbOnce sync.Once

// Model gives us a way to have a resource name its own table.
type Model interface {
	TableName() string
}

// modelTypes is ordered according to foreign key usage.
var modelTypes = []Model{
	Company{},
	Customer{},
}

func ConnectionString() string {
	cs := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		"postgres",
		"5432",
		"pavsrelic",
		"postgres",
		"reallygoodpassword",
	)
	return cs
}

func InitDB(s string) {
	dbOnce.Do(func() {
		var db *sql.DB
		var err error

		db, err = sql.Open("postgres", s)
		if err != nil {
			panic(err)
		}

		m := gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

		for _, model := range modelTypes {
			m.AddTableWithName(model, model.TableName()).SetKeys(true, "ID")
		}

		m.TypeConverter = PavsRelicTypeConverter{}

		dbmap = &m
	})
}

func GetDbMap() *gorp.DbMap {
	return dbmap
}

func NewTransaction() (*gorp.Transaction, error) {
	t, err := dbmap.Begin()
	return t, err
}

func Delete(v interface{}) (int64, error) {
	return dbmap.Delete(v)
}

// Insert is a transaction-aware helper for the built-in gorp.Insert
func Insert(v interface{}, tx *gorp.Transaction) error {
	if tx != nil {
		return tx.Insert(v)
	}
	return dbmap.Insert(v)
}

func Contains[T comparable](ss []T, v T) bool {
	for _, s := range ss {
		if s == v {
			return true
		}
	}
	return false
}

func canUpdate(colMap *gorp.ColumnMap) bool {
	return !Contains([]string{"created_at"}, colMap.ColumnName)
}

// Update is a transaction-aware helper for the built-in gorp.Update
func Update(v interface{}, tx *gorp.Transaction) (int64, error) {
	if tx != nil {
		return tx.UpdateColumns(canUpdate, v)
	}
	return dbmap.UpdateColumns(canUpdate, v)
}

// PavsRelicTypeConverter is an empty struct on top of which we implement gorp.TypeConverter
type PavsRelicTypeConverter struct{}

// ToDb implements gorp.TypeConverter
func (stc PavsRelicTypeConverter) ToDb(val interface{}) (interface{}, error) {
	return val, nil
}

// FromDb implements gorp.TypeConverter
func (stc PavsRelicTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	return gorp.CustomScanner{}, false
}

// TruncateAllTables will drop all data from the db
func TruncateAllTables() error {
	for i := len(modelTypes) - 1; i >= 0; i-- {
		query := fmt.Sprintf("DELETE FROM %s", modelTypes[i].TableName())
		_, err := dbmap.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
