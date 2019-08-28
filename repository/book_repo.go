package repository

import (
	"database/sql"
	"github.com/followme1987/bookAPI/model"
	"log"
)

type BookRepo struct{}

func (b *BookRepo) GetBooks(db *sql.DB) (books []model.Book) {
	rows, err := db.Query("select * from tblBooks")
	if err != nil {
		log.Fatal(err)
	}
	book := model.Book{}
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Year)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	return
}

func (b *BookRepo) GetBookById(db *sql.DB, id string) (book model.Book) {
	row := db.QueryRow("select * from tblBooks where id = $1", id)
	err := row.Scan(&book.Id, &book.Title, &book.Year)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (b *BookRepo) DeleteBookById(db *sql.DB, id string) (rowsAffected int64) {
	returned, err := db.Exec("delete from tblBooks where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = returned.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (b *BookRepo) UpdateBooks(db *sql.DB, book model.Book) (rowsAffected int64) {
	result, err := db.Exec("update tblBooks set title=$1,year=$2 where id=$3", &book.Title, &book.Year, &book.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (b *BookRepo) AddBooks(db *sql.DB, book model.Book) (rowsAffected int64) {
	result, err := db.Exec("insert into tblBooks (title,year) values ($1,$2)", &book.Title, &book.Year)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}
