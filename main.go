package main

import (
	"database/sql"
	"f/repository/rdbm"
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

	sdt := BuildTx(db, true)

	queries := rdbm.Queries{DB: sdt}
	bookService := usecase.NewBookService(&queries)

	_ = bookService.CheckTransaction(6)
}

// Flag true returns the Tx
func BuildTx(db *sql.DB, Flag bool) rdbm.DBTX {
	if Flag {
		tx, err := db.Begin()
		if err != nil {
			return nil
		}
		return &rdbm.SqlConnTx{DB: tx}
	} else {
		return &rdbm.SqlDBTx{DB: db}
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
