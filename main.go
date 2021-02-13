package main

import (
	"database/sql"
	"f/rdbm"
	"f/usecase"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	User     = "root"
	Password = "ducthang"
	Host     = "127.0.0.1:3306"
	Name     = "testing"
)

func main() {
	db, err := OpenConnection()
	if err != nil {
		log.Fatalf("failed with error: %v", err)
	}

	querier := rdbm.Queries{DB: db}
	store := rdbm.TxStore{Queries: &querier, DB: db}
	service := usecase.NewService(&store)

	err = service.DeleteBookOfClass(5)
	if err != nil {
		log.Fatalf("failed with error: %v", err)
	}
}

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Name,
	))

	if err != nil {
		return nil, err
	}

	return db, nil
}
