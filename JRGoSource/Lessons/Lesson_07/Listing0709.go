package main

import "fmt"

func main() {
   var taxableIncome float64 = 50000
   var taxDue float64

   if taxableIncome <= 9875 {
      taxDue = taxableIncome * 0.1
   } else if taxableIncome <= 40125 {
      taxDue = (9875 * .1) + ((taxableIncome - 9875) * .12)
   } else if taxableIncome <= 85525 {
      taxDue = (9875 * .1) + ((40125 - 9875) * .12) + ((taxableIncome - 40125) * .22)
   }

   fmt.Print("Your tax due is: ")
   fmt.Println(taxDue)
}