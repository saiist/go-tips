package src

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func redis() {
	var err error
	db, err = connectToPostgres()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Connected to PostgreSQL database")
}

func connectToPostgres() (*sql.DB, error) {
	dbHost := os.Getenv("POSTGRES_HOSTNAME")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return db, nil
}
