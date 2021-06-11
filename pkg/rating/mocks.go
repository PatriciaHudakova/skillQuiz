package rating

import (
	"errors"
	"skillQuiz/pkg/db"
)

func mockCalculateImmediateRating(answers []string) string {
	return "40"
}

func mockCalculateAverageRating(db db.IDatabase, currentRating string) (string, error) {
	return "60", nil
}

func mockCalculateAverageRatingErr(db db.IDatabase, currentRating string) (string, error) {
	return "", errors.New("test error")
}
