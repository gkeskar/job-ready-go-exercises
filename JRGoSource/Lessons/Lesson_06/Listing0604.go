package main

import (
   "fmt"
   "strconv"
)

func main() {
   var a int = 0
   for ; a < 10; {
      if (a % 2 == 0){
         fmt.Println(strconv.Itoa(a) + " is an even number")
      }
      a++ // we have to increment manually in this case
   }
}