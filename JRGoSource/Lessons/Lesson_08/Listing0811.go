package main

import (
   "fmt"
)

func main() {
   a := 4
   b := 10
   add := func() int {
      return a + b
   }
   fmt.Println(add())
}