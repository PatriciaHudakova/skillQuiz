package rating

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// PrintRatings is a wrapper function to calculate and print current & average ratings
func PrintRatings(db *sql.DB, answers []string) {
	// Based on user input, calculate the current rating
	currentRating := CalculateImmediateRating(answers)
	fmt.Printf("Your rating is: %s/100\n", currentRating)

	// Using the persisted data from the database and new user input, print the new average
	averageRating, err := CalculateAverageRating(db, currentRating)
	if err != nil {
		log.Fatalf("Something went wrong calculating your average score: %v", err)
	}
	fmt.Printf("The average rating is: %s/100", averageRating)
}

// CalculateImmediateRating calculates the rating of the current run
func CalculateImmediateRating(answers []string) string {
	var count float32
	numberOfAnswers := float32(len(answers))

	// Base case
	if numberOfAnswers <= 0 {
		return "0"
	}

	// Iterate through the responses and assign a point for each yes
	for _, param := range answers {
		// As per the task requirements, we only care about yes's, and not no's or invalid input
		if strings.EqualFold(param, "yes") {
			count++
		}
	}

	// Calculate a percentage like rating
	rating := 100 * (count/numberOfAnswers)

	// Round the rating to 0 decimal places for consistency
	return fmt.Sprintf("%.0f", rating)
}

// CalculateAverageRating calculates the average rating of all runs using the existing average in the db
func CalculateAverageRating(db *sql.DB, currentRating string) (string, error) { //TODO: create helper functions to refactor this method & test
	var average int

	// Check if table is empty
	numberOfRows, err := db.Query("SELECT * FROM averages;")
	if err != nil {
		return "", err
	}

	// If table is empty, current rating becomes the average
	if !numberOfRows.Next() {
		stmt, err := db.Prepare("INSERT INTO averages(uuid, overallAverage) values(?,?);")
		if err != nil {
			return "", err
		}

		_, err = stmt.Exec(1, currentRating)
		if err != nil {
			return "", err
		}

		return currentRating, nil
	}
	numberOfRows.Close()

	// If not, retrieve rating and calculate new rating, then add it back into the table as the new average and return
	rows, err := db.Query("SELECT overallAverage FROM averages;")
	if err != nil {
		return "", err
	}

	for rows.Next() {
		if err = rows.Scan(&average); err != nil {
			return "", err
		}
	}
	rows.Close()

	// Calculate the new average
	current, err := strconv.Atoi(currentRating)
	if err != nil {
		return "", err
	}
	newAverage := 100 * ((float64(average) + float64(current)) / 200)

	// Replace the old average with new average
	stmt, err := db.Prepare("UPDATE averages SET overallAverage=? where uuid=?") //TODO locked db
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(newAverage, 1)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.0f", newAverage), nil
}
