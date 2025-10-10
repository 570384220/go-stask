package main

import (
	"context"
	"database/sql"
	"fmt"
	"gorm/db"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	pq, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(pq)

	books, err := queries.GetBooks(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(books)
}
