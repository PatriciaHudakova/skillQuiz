package rating

import (
	"errors"
	"skillQuiz/pkg"
	"skillQuiz/pkg/db"
)

func mockCalculateImmediateRating(answers []pkg.Question) string {
	return "40"
}

func mockCalculateAverageRating(db db.IDatabase, currentRating string) (string, error) {
	return "60", nil
}

func mockCalculateAverageRatingErr(db db.IDatabase, currentRating string) (string, error) {
	return "", errors.New("test error")
}
