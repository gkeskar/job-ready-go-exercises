package main

import "fmt"

func main() {
   var a, b, c int32 = 20, 10, 8
   fmt.Println("a =", a)
   fmt.Println("b =", b)
   fmt.Println("c =", c)
   fmt.Println("a + b =", a + b)  // addition (20 + 10 = 30)
   fmt.Println("a -b =", a -b)    // subtraction (20 -10 = 10)
   fmt.Println("a * b =", a * b)  // multiplication (20 * 10 = 200)
   fmt.Println("a / b = ", a / b) // division (20 / 10 = 2)
   fmt.Println("a % c =", a % c)  // modulus (20 % 8 = 4) 

   a++ //increment by 1
   fmt.Println("a++ =", a) // a + 1 = 20 + 1 = 21 

   b-- // decrement by 1
   fmt.Println("b-- =", b) // b -1 = 10 -1 = 9
}