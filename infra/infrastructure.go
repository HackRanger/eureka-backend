package infra

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDBHandler(dbinfo string) *sql.DB {
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
