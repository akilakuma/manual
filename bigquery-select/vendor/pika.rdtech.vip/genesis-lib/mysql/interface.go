package mysql

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type SQLCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type sqlCommon struct {
	gsql gorm.SQLCommon
}

func NewSQLCommon(gsql gorm.SQLCommon) SQLCommon {
	return &sqlCommon{gsql}
}

func (sql *sqlCommon) Exec(query string, args ...interface{}) (sql.Result, error) {
	r, err := sql.gsql.Exec(query, args...)

	return r, err
}

func (sql *sqlCommon) Prepare(query string) (*sql.Stmt, error) {
	r, err := sql.gsql.Prepare(query)

	return r, err

}

func (sql *sqlCommon) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r, err := sql.gsql.Query(query, args...)

	return r, err
}

func (sql *sqlCommon) QueryRow(query string, args ...interface{}) *sql.Row {

	return sql.gsql.QueryRow(query, args...)
}
