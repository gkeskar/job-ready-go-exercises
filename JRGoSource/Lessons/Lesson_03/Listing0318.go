package main

import (
"fmt"
"strconv"
)

func main() {
   var firstNumber string
   var secondNumber string

   var firstInt int
   var secondInt int

   fmt.Print("Enter the first integer: ") // user prompt
   fmt.Scan(&firstNumber) // store input

   fmt.Print("Enter the second integer: ")
   fmt.Scan(&secondNumber)

   firstInt, _ = strconv.Atoi(firstNumber) // convert to int
   secondInt, _ = strconv.Atoi(secondNumber)

   fmt.Println(firstNumber + secondNumber) // addition of two strings
   fmt.Println(firstInt + secondInt) // addition of two ints
}