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
	defer sqlDB.Conn.Close()

	// Print and record answers to pre-defined questions
	answers := pkg.CliComponent()
	if err := rating.PrintRatings(rating.CalculateImmediateRating, rating.CalculateAverageRating, sqlDB, answers); err != nil {
		panic(err)
	}
}
