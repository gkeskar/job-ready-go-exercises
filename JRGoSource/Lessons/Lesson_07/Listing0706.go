package main

import "fmt"

func main() {
   var taxableIncome float64 = 4000
   var taxDue float64 = taxableIncome * 0.1
   fmt.Print("Your tax due is: ")
   fmt.Println(taxDue)
}