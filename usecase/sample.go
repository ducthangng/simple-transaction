package usecase

import (
	"context"
	"f/rdbm"
	"fmt"
)

type Service struct {
	repo rdbm.StoreTx
}

func NewService(repo rdbm.StoreTx) *Service {
	return &Service{repo: repo}
}

//DeleteBookOfClass deletes all the books of the chosen class.
func (s *Service) DeleteBookOfClass(cid int) error {
	ctx := context.Background()

	err := s.repo.ExecTx(ctx, func(q *rdbm.Queries) error {
		bids, err := q.QueryRelationFromDB(ctx, cid)
		if err != nil {
			return err
		}

		fmt.Println(bids)

		for _, bid := range bids {
			err = q.DeleteRelations(ctx, bid)
			if err != nil {
				return err
			}

			err = q.DeleteBook(ctx, bid)
			if err != nil {
				return err
			}

			_, err := q.QueryBookFromDB(ctx, bid)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
