package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var answer string
	var answerVaule int
	var randomNumber int

	// generate random number between 0 to 10
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	randomNumber = generator.Intn(10)
	for {
		fmt.Print("Guess what the number is: ")
		fmt.Scan(&answer)
		answerVaule, _ = strconv.Atoi(answer)
		//fmt.Println(" Random Number generated ", randomNumber)
		//fmt.Println("You guessed ", answerVaule)

		if answerVaule > randomNumber {
			fmt.Print("Too high, try again")
			continue
		} else if answerVaule < randomNumber {
			fmt.Println("Too low, try again")
			continue
		} else if answerVaule == randomNumber {
			fmt.Println("That's right, you answered the number")
			break
		}
	}
}
