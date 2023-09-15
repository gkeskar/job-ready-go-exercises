package main

import (
   "fmt"
)

func main() {
   // create an empty slice
   var s []int
   fmt.Println(s)
   s = append(s, 10)
   fmt.Println(s)
   s = append(s, 11)
   fmt.Println(s)
}