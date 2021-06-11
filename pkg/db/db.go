package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DriverName     = "sqlite3"
	DataSourceName = "identifier.sqlite"
)

type Database struct {
	Conn *sql.DB
}

// InitDB initialises and creates a connection to our database
func InitDB() (*Database, error) {
	// Initialise the database
	db, err := sql.Open(DriverName, DataSourceName)
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

	return &Database{Conn: db}, nil
}

// GetAllRows retrieves all rows from the averages table
func (db *Database) GetAllRows() (*sql.Rows, error) {
	numberOfRows, err := db.Conn.Query("SELECT * FROM averages;")
	if err != nil {
		return nil, err
	}

	return numberOfRows, nil
}

// MakeCurrentRatingTheAverage makes the current run's rating the new overall average
func (db *Database) MakeCurrentRatingTheAverage(currentRating string) error {
	stmt, err := db.Conn.Prepare("INSERT INTO averages(uuid, overallAverage) values(?,?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(1, currentRating)
	if err != nil {
		return err
	}

	return nil
}

// GetOverallAverageFromDB retrieves the stored overall average from db
func (db *Database) GetOverallAverageFromDB() (int, error) {
	var average int

	rows, err := db.Conn.Query("SELECT overallAverage FROM averages;")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	for rows.Next() {
		if err = rows.Scan(&average); err != nil {
			log.Fatal(err)
			return 0, err
		}
	}
	rows.Close()

	return average, nil
}

// UpdateAverage updates the table with new average
func (db *Database) UpdateAverage(newAverage int) error {
	// Replace the old average with new average
	stmt, err := db.Conn.Prepare("UPDATE averages SET overallAverage=? where uuid=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newAverage, 1)
	if err != nil {
		return err
	}

	return nil
}
