package rating

import (
	"errors"
	"skillQuiz/pkg/db"
)

func mockCalculateImmediateRating(answers []string) string {
	return "40"
}

func mockCalculateAverageRating(db *db.Database, currentRating string) (string, error) {
	return "60", nil
}

func mockCalculateAverageRatingErr(db *db.Database, currentRating string) (string, error) {
	return "", errors.New("test error")
}
