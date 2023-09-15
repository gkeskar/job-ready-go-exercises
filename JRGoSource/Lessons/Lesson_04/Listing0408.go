package main

import "fmt"

func main() {
   var a, b int8 = 100, 70 // int8 cannot be larger than 128

   fmt.Println("a =", a)
   fmt.Println("b = ", b)
   fmt.Println("a == b:", a == b) //checks for equality
   fmt.Println("a != b:", a != b) //checks for inequality
   fmt.Println("a > b:", a > b)
   fmt.Println("a < b:", a < b)
   fmt.Println("a >= b:", a >= b)
}