package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func AskQuestions() []string {
	// Initialise a new scanner that can scan a line from the console and hold the input as a variable
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Skill Questionnaire")

	fmt.Println("-------------------")
	fmt.Println("Can you code in Ruby?")
	scanner.Scan()
	ruby := scanner.Text()
	fmt.Println("Can you code in JavaScript?")
	scanner.Scan()
	js := scanner.Text()
	fmt.Println("Can you code in Swift (iOS)?")
	scanner.Scan()
	swift := scanner.Text()
	fmt.Println("Can you code in Java (Android)?")
	scanner.Scan()
	java := scanner.Text()
	fmt.Println("Can you code in C#?")
	scanner.Scan()
	cs := scanner.Text()
	fmt.Println("-------------------")

	return []string{ruby, js, swift, java, cs}
}
