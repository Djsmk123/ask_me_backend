package db

import (
	"context"
	"database/sql"
	"fmt"
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

func (store *SQLCExec) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb erro: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
