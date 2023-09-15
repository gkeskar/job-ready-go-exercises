package main

import (
   "fmt"
)

// add takes as input a and b of type int and returns an int
func add(a int, b int) int {
   return a + b
}

// add2 is a short version of add
func add2(a, b int) int {
   return a + b
}

// multiply takes as input a and b of type int and returns an int
func multiply(a int, b int) int {
   return a * b
}

func main() {
   c := 5
   d := 6
   fmt.Println("c =", c)
   fmt.Println("d =", d)
   fmt.Println("add result:", add(c, d))
   fmt.Println("add2 result:", add2(c, d))
   fmt.Println("multiply result:", multiply(c, d))
}