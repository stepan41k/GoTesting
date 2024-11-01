package main

import (
	"fmt"
	"log"


	"github.com/stepan41k/lessonPostgres/pkg/repository"
)

type book struct{
	id int
	name string
	price int
	authorID int
	genreID int
}

const connStr = "postgres://postgres:admin@localhost:5432/cources"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	item, err := db.GetBookByID(8)
	if err != nil {
		log.Fatal(err.Error())
	}


	fmt.Printf("%d: %s, price: %d, author id: %d, genre id: %d\n", item.ID, item.Name, item.Price, item.AuthorID, item.GenreID)


	// book := models.Book{GenreID: 1, AuthorID: 1, Name: "Идиот", Price: 150}
	// err = db.NewBook(book)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }




}