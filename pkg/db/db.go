package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"
	dataSourceName = "identifier.sqlite"
)

func InitDB() (*sql.DB, error) {
	// Initialise the database
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
		return nil, err
	}
	log.Println("Successfully connected to the database!")

	// Test DB is ready to accept connections
	if err := db.Ping(); err != nil {
		log.Fatalf("DB is unable to accept connections: %v", err)
		return nil, err
	}
	log.Println("Ready to accept connections!")

	return db, nil
}