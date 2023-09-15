package main

import "fmt"

func main() {
   // ask user for the gross income
   fmt.Print("Enter your gross income from your W-2 for 2020:")

   var grossIncome float64
   fmt.Scanln(&grossIncome) // take input from user
   fmt.Print("Your gross income is: ")
   fmt.Println(grossIncome)
}