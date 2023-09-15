package main

import "fmt"

func main() {
	var questions [10]string
	questions[0] = "What is your name?"
	questions[1] = "Where do you leave?"
	questions[2] = "When is your birthday?"
	questions[3] = "What is your favorite hobby?"
	questions[4] = "What is your favorite food?"

	//var question string
	var answers [10]string
	var answer string

	fmt.Println(questions[0])
	fmt.Scan(&answer)
	answers[0] = answer

	fmt.Println(questions[1])
	fmt.Scan(&answer)
	answers[1] = answer

	fmt.Println(answers)
}
