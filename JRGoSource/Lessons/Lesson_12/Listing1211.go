package main

import (
   "fmt"
)

func main() {
   numbers := [7]int{0,1,2,5,798,43,78}
   fmt.Println(numbers)

   s := numbers[0:4] 
   fmt.Println(s)

   for i, v := range s {
      fmt.Println("Element", i, "is", v)
   }
}