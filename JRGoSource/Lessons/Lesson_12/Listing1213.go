package main

import (
   "fmt"
)

func main() {
   var s = []int{10, 11} 
   fmt.Println("Slice s: ",s)

   // create a destination slice
   c := make([]int, len(s))

   // copy everything in s to c
   num := copy(c, s) // returns the minimum number of elements in the slices
   fmt.Println("Number of elements copied:", num)
   fmt.Println("Slice c:", c)
}