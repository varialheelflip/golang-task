package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Book struct {
	Id     uint
	Title  string
	Author string
	Price  float64
}

func getBooks(db *sqlx.DB) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books where price >= ?", 50)
	return books, err
}

func main() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(192.168.200.130:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	books, err := getBooks(db)
	fmt.Println(books)
	fmt.Println(err)
}
