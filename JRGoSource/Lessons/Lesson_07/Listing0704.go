package main

import "fmt"

func main() {

   // ask user for the gross income
   fmt.Print("Enter your gross income from your W-2 for 2020:")

   var grossIncome float64
   fmt.Scanln(&grossIncome) // take gross income input from user
   fmt.Print("Your gross income is: ")
   fmt.Println(grossIncome)

   fmt.Print("How many dependents are you claiming? ")
   var numDep int
   fmt.Scanln(&numDep) // take number of dependents input from user
   fmt.Print("Your claimed number of dependents is: ")
   fmt.Println(numDep)

   //calculate taxable income

   var taxableIncome float64
   taxableIncome = grossIncome - 12200 - (2000 * float64(numDep))
   fmt.Print("Your taxable income is: ")
   fmt.Println(taxableIncome)
}