package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	Text   string
	Answer string
	Value  int
}

func createQuestions() []Question {
	questions := []Question{{
		Text:   "Can you code in Ruby?",
		Answer: "",
		Value:  0,
	}, {
		Text:   "Can you code in JavaScript?",
		Answer: "",
		Value:  0,
	}, {
		Text:   "Can you code in Swift (iOS)?",
		Answer: "",
		Value:  0,
	}, {
		Text:   "Can you code in Java (Android)?",
		Answer: "",
		Value:  0,
	}, {
		Text:   "Can you code in C#?",
		Answer: "",
		Value:  0,
	},
	}

	return questions
}

// CliComponent outputs and records answers to pre-defined questions
func CliComponent() []Question {
	// Initialise a new scanner that can scan a line from the console and hold the input as a variable
	scanner := bufio.NewScanner(os.Stdin)
	questions := createQuestions()
	var populatedQs []Question

	fmt.Println("Skill Questionnaire")
	fmt.Println("-------------------")
	// Iterate through questions, check the response and assign a score
	for _, question := range questions {
		// Print question and record response
		fmt.Println(question.Text)
		scanner.Scan()
		answer := scanner.Text()

		// Populate the remainder of the struct with correct values
		question.Answer = answer
		switch strings.ToLower(answer) {
		case "yes":
			question.Value = 1
		case "y":
			question.Value = 1
		case "no":
			question.Value = 0
		case "n":
			question.Value = 0
		default:
			question.Value = 0
		}

		// Populate the the questions array with newly populated structs
		populatedQs = append(populatedQs, question)
	}
	fmt.Println("-------------------")

	return populatedQs
}
