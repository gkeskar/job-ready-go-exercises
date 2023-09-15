package main

import (
   "fmt"
)

func main() {
   // define an array with 7 elements
   numbers := [7]int{0,1,2,5,798,43,78}
   fmt.Println(numbers)

   // define a slice mySlice based on numbers in the array
   mySlice := numbers[1:5]
   fmt.Println(mySlice)

   fmt.Println("Length of slice:", len(mySlice))

   fmt.Println("Slice capacity:", cap(mySlice))
}