package rdbm

import (
	"context"
)

type Queries struct {
	DB DBTX
}

func (q *Queries) QueryBookFromDB(ctx context.Context, bid int) (string, error) {
	var s string
	err := q.DB.QueryRow("select bookname from books where bid = ?", bid).Scan(&s)

	if err != nil {
		return s, err
	}

	return s, nil
}

func (q *Queries) QueryRelationFromDB(ctx context.Context, cid int) ([]int, error) {
	var rest []int
	rows, err := q.DB.Query("select bid from bookclassrelations where cid = ?", cid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bid int
		err := rows.Scan(&bid)
		if err != nil {
			return nil, err
		}

		rest = append(rest, bid)
	}

	return rest, nil
}

func (q *Queries) DeleteBook(ctx context.Context, bid int) error {
	_, err := q.DB.Exec("delete from books where bid = ?", bid)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) DeleteRelations(ctx context.Context, bid int) error {
	_, err := q.DB.Exec("delete from bookclassrelations where bid = ?", bid)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) EnableTx(txFunc func() error) error {
	return q.DB.TxEnd(txFunc)
}
