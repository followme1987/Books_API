package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/followme1987/bookAPI/model"
	"github.com/followme1987/bookAPI/repository"
	"github.com/followme1987/bookAPI/util"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct{}

func (c *Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := repository.BookRepo{}
		books := repo.GetBooks(db)
		w.Header().Set("Content-Type", "application/json")
		util.SendMsg(w, books)
	}
}

func (c *Controller) GetBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		repo := repository.BookRepo{}
		book := repo.GetBookById(db, id)
		w.Header().Set("Content-Type", "application/json")
		util.SendMsg(w, book)
	}
}

func (c *Controller) DeleteBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		repo := repository.BookRepo{}
		rowsAffected := repo.DeleteBookById(db, id)
		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}

func (c *Controller) UpdateBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := model.Book{}
		json.NewDecoder(r.Body).Decode(&book)
		repo := repository.BookRepo{}
		rowsAffected := repo.UpdateBooks(db, book)
		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}

func (c *Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := model.Book{}
		json.NewDecoder(r.Body).Decode(&book)
		repo := repository.BookRepo{}
		rowsAffected := repo.AddBooks(db, book)
		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}
