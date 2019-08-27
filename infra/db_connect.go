package infra

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetDb() (db *sql.DB) {
	conn_str := os.Getenv("ELEPHANTSQL_CONN")
	db, err := sql.Open("postgres", conn_str)
	if err != nil {
		log.Fatal(err)
	}
	return
}
