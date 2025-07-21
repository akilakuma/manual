package infra

import "database/sql"

func NewDB() (*sql.DB, error) {

	return sql.Open("mysql", "user:pass@tcp(localhost:3306)/mydb")
}
