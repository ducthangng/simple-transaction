package usecase

import (
	"context"
	"f/repository"
)

type BookService struct {
	BookDataInterface repository.BookDataInterface
}

func NewBookService(repo repository.BookDataInterface) *BookService {
	return &BookService{BookDataInterface: repo}
}

//DeleteBookOfClass deletes all the books of the chosen class.
func (bs *BookService) CheckTransaction(cid int) error {
	ctx := context.Background()
	bdi := bs.BookDataInterface

	err := bdi.EnableTx(func() error {
		bids, err := bdi.QueryRelationFromDB(ctx, cid)
		if err != nil {
			return err
		}

		for _, bid := range bids {
			_, err := bdi.QueryBookFromDB(ctx, bid)
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
