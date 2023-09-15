package main

import "fmt"

func main() {
   var tax float32 = 0.065 // IEEE-754 32-bit floating point number
   var i float64 = 0.000006 // IEEE-754 64-bit float
   var cnumber complex64 = 1 + 4i //a complex number with float 32
   //real and imaginary numbers

   fmt.Println("tax float32:", tax)
   fmt.Println("i float64:", i)
   fmt.Println("cnumber complex64:", cnumber)
}