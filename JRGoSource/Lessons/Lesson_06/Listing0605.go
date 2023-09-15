package main

import (
   "fmt"
   "strconv"
)

func main() {
   var a int = 0
   for a < 10 { // remove the semicolons here
      if (a % 2 == 0){
         fmt.Println(strconv.Itoa(a) + " is an even number")
      }

      a++
   }
}