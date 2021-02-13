package rdbm

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	DB DBTX
}

type Book struct {
	BID      int
	BookName string
	Info     string
}

type BookClass struct {
	BID int
	CID int
}

type TxStore struct {
	DB *sql.DB
	*Queries
	// StudentQueries
	// TeacherQueries
	// AdminQueries
}

// ExecTx executes a function within a database transaction
func (store *TxStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := &Queries{DB: tx}

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
