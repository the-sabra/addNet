package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// Initialize the database connection and perform any necessary setup.
	// This is a placeholder function and should be implemented according to your database setup.
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
