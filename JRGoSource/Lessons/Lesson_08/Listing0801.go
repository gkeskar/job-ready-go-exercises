package main

import (
   "fmt"
)

// add takes as input a and b of type int and returns an int
func add(a int, b int) int {
   return a + b
}

func main() {
   fmt.Println("add function results:", add(4, 6))
}