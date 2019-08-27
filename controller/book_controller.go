package controller

import (
	"database/sql"
	"github.com/followme1987/bookAPI/model"
	"github.com/followme1987/bookAPI/util"
	"log"
	"net/http"
)

type Controller struct{}


func (c *Controller) GetBooks(db *sql.DB)  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		rows, err := db.Query("select * from tblBooks")
		if err != nil {
			log.Fatal(err)
		}

		books := []model.Book{}
		book := model.Book{}
		for rows.Next(){
			err := rows.Scan(&book.Id,&book.Title,&book.Year)
			if err != nil{
				log.Fatal(err)
			}
			books = append(books,book)
		}
		util.SendMsg(w,books)
	}
}
