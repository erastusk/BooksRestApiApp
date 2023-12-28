package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var constr = "postgres://postgres:@localhost:5432/postgres?sslmode=disable"

func SqlConn() *sql.DB {
	db, err := sql.Open("postgres", constr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
