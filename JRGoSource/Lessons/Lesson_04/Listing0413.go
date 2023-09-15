package main

import (
   "fmt"
   "math"
)

func main() {
   var a, b, c float64 = -30.5, 45.6, 4
   fmt.Println("a =", a, "\nb =", b, "\nc =", c)

   fmt.Println("math.Abs(a) =", math.Abs(a)) // \|30.5\|
   fmt.Println("math.Ceil(b) =", math.Ceil(b)) // 46
   fmt.Println("math.Floor(b) =", math.Floor(b)) //45
   fmt.Println("math.Exp(a) =", math.Exp(a)) // exponential of 46.6
   fmt.Println("math.Sqrt(c) =", math.Sqrt(c)) // 2
   fmt.Println("math.Trunc(a) =", math.Trunc(a)) // -30
   fmt.Println("math.Round(a) =", math.Round(a)) // -31
   fmt.Println("math.Pow(b,c) =", math.Pow(b,c)) // b to the power of c
}