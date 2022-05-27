package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Record struct {
	*Queries
	db *sql.DB
}

func NewRecord(db *sql.DB) *Record {
	return &Record{
		db:      db,
		Queries: New(db),
	}
}

func (record *Record) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := record.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Tx Err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
