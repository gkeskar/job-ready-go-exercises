package main

import (
   "fmt"
)

func main() {
   numbers := [7]int{0, 1, 2, 5, 798, 43, 78}
   fmt.Println("Numbers array:", numbers)

   // define a slice with the first 4 elements
   slice_1 := numbers[0:4]
   fmt.Println("slice_1:", slice_1)

   // define slice from the second element through the end of the array
   slice_2 := numbers[1:]
   fmt.Println("slice_2:", slice_2)

   // define a slice with the first 5 elements
   slice_3 := numbers[:5]
   fmt.Println("slice_3:", slice_3)
}