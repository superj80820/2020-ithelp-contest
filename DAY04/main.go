package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	HOST     = "db"
	DATABASE = "postgres"
	USER     = "user"
	PASSWORD = "mysecretpassword"
)

func main() {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE),
	)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully created connection to database")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!")
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
