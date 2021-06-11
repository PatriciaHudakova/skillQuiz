//go:generate mockgen -source=interface.go -destination=mock.go -package db
package db

import "database/sql"

type IDatabase interface {
	GetAllRows() (*sql.Rows, error)
	MakeCurrentRatingTheAverage(currentRating string) error
	GetOverallAverageFromDB() (int, error)
	UpdateAverage(newAverage int) error
	IsEmpty(rows *sql.Rows) bool
}
