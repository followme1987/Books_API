package main

import (
	"github.com/followme1987/bookAPI/controller"
	"github.com/followme1987/bookAPI/infra"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)
func init(){
	gotenv.Load()
}

func main() {
	db := infra.GetDb()

	ctl := controller.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/books", ctl.GetBooks(db)).Methods("GET")

	err := http.ListenAndServe(":50000", router)
	if err != nil {
		log.Fatal(err)
	}
}
