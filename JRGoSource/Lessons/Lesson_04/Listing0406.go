package main

import "fmt"

func main() {
   var a, b, c, d int = 100, 50, 25, 4
   fmt.Println("a =", a, "\nb =", b, "\nc =", c, "\nd = ", d)

   a *= b // a * b = 100 * 50 = 5000 (a = 5000)
   fmt.Println("a *= b =", a)

   b /= c // b / c = 50 / 25 = 2 (b = 2)
   fmt.Println("b /= c =", b)

   c %= d // c % d = 25 % 4 = 1 (c = 1)
   fmt.Println("c %= d =", c)

   fmt.Println("d =", d)
}