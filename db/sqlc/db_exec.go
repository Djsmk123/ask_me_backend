package db

import (
	"database/sql"
)

type DBExec interface {
	Querier
}
type SQLCExec struct {
	*Queries
	db *sql.DB
}

func NewDBSQLExec(db *sql.DB) DBExec {
	return &SQLCExec{
		db:      db,
		Queries: New(db),
	}
}
