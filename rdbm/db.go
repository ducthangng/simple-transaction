package rdbm

import "context"

type DBrepo interface {
	QueryBookFromDB(ctx context.Context, id int) (string, error)
	QueryRelationFromDB(ctx context.Context, cid int) ([]int, error)
	DeleteBook(ctx context.Context, bid int) error
	DeleteRelations(ctx context.Context, bid int) error
}

type StoreTx interface {
	DBrepo

	//transaction handlers
	ExecTx(ctx context.Context, fn func(*Queries) error) error
}
