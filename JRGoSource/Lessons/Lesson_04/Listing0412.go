package main

import "fmt"

func main() {
   var a bool = true
   fmt.Println("a =", a)
   var b bool = false
   fmt.Println("b =", b)

   fmt.Println("a && b =", a && b) // AND operator
   fmt.Println("a || b =", a || b) // OR operator
   fmt.Println("!a =", !a) // NOT operator
}