package main

import (
   "fmt"
)

func main() {
   s := make([]int, 10)
   fmt.Println("Original slice:", s)

   s[0]=99
   fmt.Println("Updated slice:", s)
}