package main

import (
	"skillQuiz/pkg"
	"skillQuiz/pkg/db"
	"skillQuiz/pkg/rating"
)

func main() {
	// Connect to the database
	sqlDB, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Ask questions, record answer and print ratings
	answers := pkg.AskQuestions()
	if err := rating.PrintRatings(rating.CalculateImmediateRating, rating.CalculateAverageRating, sqlDB, answers); err != nil {
		panic(err)
	}
}
