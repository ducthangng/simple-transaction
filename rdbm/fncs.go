package rdbm

import (
	"context"
)

func (q *Queries) QueryBookFromDB(ctx context.Context, bid int) (string, error) {
	var s string
	err := q.DB.QueryRowContext(ctx, "select bookname from books where bid = ?", bid).Scan(&s)

	if err != nil {
		return s, err
	}

	return s, nil
}

func (q *Queries) QueryRelationFromDB(ctx context.Context, cid int) ([]int, error) {
	var rest []int
	rows, err := q.DB.QueryContext(ctx, "select bid from bookclassrelations where cid = ?", cid)

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
	_, err := q.DB.ExecContext(ctx, "delete from books where bid = ?", bid)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) DeleteRelations(ctx context.Context, bid int) error {
	_, err := q.DB.ExecContext(ctx, "delete from bookclassrelations where bid = ?", bid)
	if err != nil {
		return err
	}

	return nil
}

/*
func (store *SQLStore) TransferTx(ctx context.Context, id int) error {
	err := store.execTx(ctx, func(q *Queries) error {
		fmt.Println("3")
		result, err := q.QueryRelationFromDB(ctx, id)
		if err != nil {
			return err
		}

		err = q.deleteRelations(ctx, result)
		if err != nil {
			return err
		}

		err = q.deleteBook(ctx, 6)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
*/
