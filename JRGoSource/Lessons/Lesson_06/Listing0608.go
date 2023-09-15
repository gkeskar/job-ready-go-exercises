package main

import "fmt"

func main() {
   var message string = "HELLO WORLD"

   fmt.Println(message)

   for idx := 0; idx < len(message); idx ++ {
      fmt.Println(string(message[idx]))
   }
}