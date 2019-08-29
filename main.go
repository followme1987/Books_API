package main

import (
	"github.com/followme1987/bookAPI/controller"
	"github.com/followme1987/bookAPI/infra"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

func init() {
	gotenv.Load()
}

func main() {
	db := infra.GetDb()

	ctl := controller.Controller{}

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	router.HandleFunc("/books", ctl.GetBooks(db)).Methods("GET")
	router.HandleFunc("/book/{id}", ctl.GetBookById(db)).Methods("GET")
	router.HandleFunc("/book/{id}", ctl.DeleteBookById(db)).Methods("DELETE")
	router.HandleFunc("/books", ctl.UpdateBooks(db)).Methods("PUT")
	router.HandleFunc("/book", ctl.AddBook(db)).Methods("POST")

	err := http.ListenAndServe(":50000", router)
	if err != nil {
		log.Fatal(err)
	}
}
