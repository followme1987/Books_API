package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/followme1987/bookAPI/model"
	"github.com/followme1987/bookAPI/repository"
	"github.com/followme1987/bookAPI/util"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
)

type Controller struct{}

func (c *Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var books []model.Book
		if db != nil {
			repo := repository.BookRepo{}
			books = repo.GetBooks(db)
		}
		w.Header().Set("Content-Type", "application/json")
		util.SendMsg(w, books)
	}
}

func (c *Controller) GetBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var uriRegex = regexp.MustCompile(`/book/\d`)
		if uriRegex.MatchString(r.RequestURI) {
			id := mux.Vars(r)["id"]
			book := model.Book{}
			if db != nil {
				repo := repository.BookRepo{}
				book = repo.GetBookById(db, id)
			}
			w.Header().Set("Content-Type", "application/json")
			util.SendMsg(w, book)
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c *Controller) DeleteBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var uriRegex = regexp.MustCompile(`/book/\d`)
		if uriRegex.MatchString(r.RequestURI) {
			id := mux.Vars(r)["id"]
			var rowsAffected int64 = 0
			if db != nil {
				repo := repository.BookRepo{}
				rowsAffected = repo.DeleteBookById(db, id)
			}
			w.Header().Set("Content-Type", "text/plain")
			util.SendMsg(w, rowsAffected)
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c *Controller) UpdateBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		book := model.Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var rowsAffected int64 = 0
		if db != nil {
			repo := repository.BookRepo{}
			rowsAffected = repo.UpdateBooks(db, book)
		}
		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)

	}
}

func (c *Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		book := model.Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		var rowsAffected int64 = 0
		if db != nil {
			repo := repository.BookRepo{}
			rowsAffected = repo.AddBooks(db, book)
		}
		w.Header().Set("Content-Type", "text/plain")
		util.SendMsg(w, rowsAffected)
	}
}
