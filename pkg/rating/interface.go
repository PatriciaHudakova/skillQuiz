//go:generate mockgen -source=interface.go -destination=mock.go -package rating

package rating

import "database/sql"

// Rating is a rating interface
type Rating interface {
	PrintRatings(db *sql.DB, answers []string)
	CalculateImmediateRating(answers []string) string
	CalculateAverageRating(db *sql.DB, currentRating string) (string, error)
}
