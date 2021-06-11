package pkg

import (
	"bufio"
	"fmt"
	"os"
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
	populated := []Question{}

	fmt.Println("Skill Questionnaire")

	fmt.Println("-------------------")
	for _, question := range questions {
		fmt.Println(question.Text)
		scanner.Scan()
		answer := scanner.Text()
		question.Answer = answer
		switch answer {
		case "yes":
			question.Value = 1
		case "no":
			question.Value = 0
		default:
			question.Value = 0
		}
		populated = append(populated, question)
	}
	fmt.Println("-------------------")

	fmt.Println(populated)
	return populated
}
