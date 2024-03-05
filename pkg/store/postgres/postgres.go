package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DBconnection(username, password, localhost, port, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", username, password, localhost, port, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error: Unable to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error: Unable to ping database: %v", err)
	}

	return db, nil
}