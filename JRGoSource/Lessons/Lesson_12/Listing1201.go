package main

import (
   "fmt"
   "reflect"
)

func main() {
   // define an array with 7 elements
   numbers := [7]int{0,1,2,5,798,43,78}
   fmt.Println("array value:", numbers)
   fmt.Println("array type:", reflect.TypeOf(numbers))

   // define a slice s based on the numbers array
   s := numbers[0:4]
   fmt.Println("slice value:", s)
   fmt.Println("slice type:", reflect.TypeOf(s))
}