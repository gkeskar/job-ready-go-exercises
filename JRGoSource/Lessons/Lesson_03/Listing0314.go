package main

import "fmt"

func main() {
   var myVar string = "Hello, World!"
   var myAge int = 99

   fmt.Println(myVar)
   fmt.Println(&myVar)
   fmt.Println(myAge)
   fmt.Println(&myAge)
}