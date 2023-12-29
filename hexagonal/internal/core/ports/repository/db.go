package ports

import (
	"database/sql"
)

type Database interface {
	GetDB() *sql.DB
}
