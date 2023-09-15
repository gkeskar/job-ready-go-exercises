package main

import "fmt"

func main() {
   var a, b, c int = 100, 70, 50
   fmt.Println("a =", a, "\nb =", b, "\nc =", c)

   a += b // a + b = 70 + 100 = 170 (a = 170)
   fmt.Println("a += b =", a)

   b -= c // b -c = 90 - 50 = 20 (b = 20)
   fmt.Println("b -= c =", b)

   fmt.Println("\nc =", c)
}