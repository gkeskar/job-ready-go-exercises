package main

import (
   "fmt"
   "reflect"
)

func main() {
   var s[]int // define a slice
   fmt.Println(s) // display an empty slice
   fmt.Println("Slice type:", reflect.TypeOf(s)) // type of the slice
}
