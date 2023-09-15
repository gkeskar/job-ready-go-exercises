package main

import "fmt"

func main() {
   var firstNumber string
   var secondNumber string

   fmt.Print("Enter the first integer: ") // user prompt
   fmt.Scan(&firstNumber) // store input

   fmt.Print("Enter the second integer: ")
   fmt.Scan(&secondNumber)

   fmt.Println(firstNumber + secondNumber) // addition of two strings
}