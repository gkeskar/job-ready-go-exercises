package main

import "fmt"

func main() {
   var taxableIncome float64 = 140000
   var max10 float64 = 9875
   var max12 float64 = 40125
   var max22 float64 = 85525
   var max24 float64 = 163300
   //var max32 float64 = 207350
   //var max35 float64 = 518400

   var tier10_tax float64 = max10 * .1
   var tier12_tax float64 = tier10_tax + ((max12 - max10) * .12)
   var tier22_tax float64 = tier12_tax + ((max22 - max12) * .22)
   //var tier24_tax float64 = tier22_tax + ((max24 - max22) * .24)
   //var tier32_tax float64 = tier24_tax + ((max32 - max24) * .32)
   //var tier35_tax float64 = tier32_tax + ((max35 - max32) * .35)

   var taxDue float64

   if taxableIncome <= max10 {
      taxDue = taxableIncome * 0.1
   } else if taxableIncome <= max12 {
      taxDue = tier10_tax + ((taxableIncome - max10) * .12)
   } else if taxableIncome <= max22 {
      taxDue = tier12_tax + ((taxableIncome - max12) * .22)
   } else if taxableIncome <= max24 {
      taxDue = tier22_tax + ((taxableIncome - max24) * .32)
   }

   fmt.Print("Your tax due is: ")
   fmt.Println(taxDue)
}