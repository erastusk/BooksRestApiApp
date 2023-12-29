package postgres

import (
	"database/sql"
	"log"
)

type PostgresRepository struct {
	dbconstring string
}

func NewRepository() *PostgresRepository {
	return &PostgresRepository{
		dbconstring: "postgres://postgres:@localhost:5432/postgres?sslmode=disable",
	}
}

func (p *PostgresRepository) GetDB() *sql.DB {
	db, err := sql.Open("postgres", p.dbconstring)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
