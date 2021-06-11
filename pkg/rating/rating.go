package rating

import (
	"fmt"
	"skillQuiz/pkg/db"
	"strconv"
	"strings"
)

type CurrentRun func(answers []string) string
type AverageRun func(db *db.Database, currentRating string) (string, error)

// PrintRatings is a wrapper function to calculate and print current & average ratings
func PrintRatings(currentRunFunc CurrentRun, averageRunFunc AverageRun, db *db.Database, answers []string) error {
	// Based on user input, calculate the current rating
	currentRating := currentRunFunc(answers)
	fmt.Printf("Your rating is: %s/100\n", currentRating)

	// Using the persisted data from the database and new user input, print the new average
	averageRating, err := averageRunFunc(db, currentRating)
	if err != nil {
		return fmt.Errorf("something went wrong calculating your average score: %v", err)
	}
	fmt.Printf("The average rating is: %s/100", averageRating)

	return nil
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
	rating := 100 * (count / numberOfAnswers)

	// Round the rating to 0 decimal places for consistency
	return fmt.Sprintf("%.0f", rating)
}

// CalculateAverageRating calculates the average rating of all runs using the existing average in the db
func CalculateAverageRating(db *db.Database, currentRating string) (string, error) {
	var average int

	// Retrieve all rows from the averages table
	numberOfRows, err := db.GetAllRows()
	if err != nil {
		return "", err
	}

	// If there are no entries, current rating becomes the average
	if !numberOfRows.Next() {
		if err := db.MakeCurrentRatingTheAverage(currentRating); err != nil {
			return "", fmt.Errorf("unable to persist current average: %v", err)
		}
		numberOfRows.Close()
		return currentRating, nil
	}
	numberOfRows.Close()

	// If not, retrieve rating and calculate new rating, then add it back into the table as the new average and return
	average, err = db.GetOverallAverageFromDB()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve average: %v", err)
	}

	// Calculate the new average
	current, err := strconv.Atoi(currentRating)
	if err != nil {
		return "", err
	}

	// Convert to float as using an integer would result in 0
	newAverage := 100 * ((float64(average) + float64(current)) / 200)

	// Replace the old average with new average
	if err = db.UpdateAverage(int(newAverage)); err != nil {
		return "", fmt.Errorf("unable to update average: %v", err)
	}

	return fmt.Sprintf("%.0f", newAverage), nil
}
