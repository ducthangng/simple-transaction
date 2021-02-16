package repository

import (
	"context"
	"f/txdataservice"
)

type BookDataInterface interface {
	QueryBookFromDB(ctx context.Context, id int) (string, error)
	QueryRelationFromDB(ctx context.Context, cid int) ([]int, error)
	DeleteBook(ctx context.Context, bid int) error
	DeleteRelations(ctx context.Context, bid int) error
	txdataservice.TxDataInterface
}
