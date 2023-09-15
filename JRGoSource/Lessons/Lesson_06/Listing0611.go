package main

import (
   "fmt"
   "strconv"
)

func main() {
   for ctr := 0; ctr < 10; ctr ++{
      if (ctr % 2 == 0){
         continue // continue to next iteration; i.e., ignore even values
      }
      fmt.Println(strconv.Itoa(ctr) + " is an odd number")
      }
}