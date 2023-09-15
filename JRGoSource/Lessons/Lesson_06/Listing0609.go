package main

import "fmt"

func main() {
   var message string = "HELLO WORLD"
   fmt.Println(message)

   for idx, c := range message {
      fmt.Println(idx) //index
      fmt.Println(string(c)) //value
   }
}