package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// Initialize the database connection and perform any necessary setup.
	// This is a placeholder function and should be implemented according to your database setup.
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))

	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(30)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
