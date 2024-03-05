package main

import (
	"log"
	// "net/http"

	"go-web-robotek/pkg/store/postgres"
)

func main() {
	db, err := postgres.DBconnection("postgres", "1234", "localhost", "5432", "robotek")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

	sqlStatement := `
INSERT INTO teachers (fullName, email, password, phoneNumber)
VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, "test test", "test@edu.kz", "test", "123123")
	if err != nil {
		panic(err)
	}

	// mux := http.NewServeMux()

	// log.Print("starting server on :4000")

	// err = http.ListenAndServe(":4000", mux)
	// log.Fatal(err)

}
