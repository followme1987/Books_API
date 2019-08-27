package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/followme1987/bookAPI/model"
	"github.com/followme1987/bookAPI/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Controller struct{}

func (c *Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("select * from tblBooks")
		if err != nil {
			log.Fatal(err)
		}

		var books []model.Book
		book := model.Book{}
		for rows.Next() {
			err := rows.Scan(&book.Id, &book.Title, &book.Year)
			if err != nil {
				log.Fatal(err)
			}
			books = append(books, book)
		}
		
		w.Header().Set("Content-Type", "application/json")
		util.SendMsg(w, books)
	}
}

func (c *Controller) GetBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		row := db.QueryRow("select * from tblBooks where id = $1", id)
		var book model.Book

		err := row.Scan(&book.Id, &book.Title, &book.Year)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		util.SendMsg(w, book)
	}
}

func (c *Controller) DeleteBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		returned, err := db.Exec("delete from tblBooks where id = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		rowsAffected, err := returned.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}

func (c *Controller) UpdateBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		book := model.Book{}
		json.NewDecoder(r.Body).Decode(&book)

		result, err := db.Exec("update tblBooks set title=$1,year=$2 where id=$3", &book.Title, &book.Year, &book.Id)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/plain")

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}

func (c *Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := model.Book{}
		json.NewDecoder(r.Body).Decode(&book)

		result, err := db.Exec("insert into tblBooks (title,year) values ($1,$2)", &book.Title, &book.Year)
		if err != nil {
			log.Fatal(err)
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}
