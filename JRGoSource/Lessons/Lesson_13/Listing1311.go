package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int)
   m["Hi"] = 20
   m["How"] = 245
   m["hi"] = 23
   m["hello"] = 2
   m["hey"] = 4
   m["weather"] = 2 
   m["greet"] = 35

   for key, value := range m {
      fmt.Println("Key:", key, "Value:", value)
   }
}
