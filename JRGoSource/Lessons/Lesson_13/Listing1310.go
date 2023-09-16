package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int)
   m["Hi"] = 20
   m["How"] = 245
   fmt.Println("Map value:", m)

   // check to see if "Hi" exists
   val, ok := m["Hi"]
   fmt.Println("Value of val:", val)
   fmt.Println("Value of ok :", ok)

   if ok == true{
      fmt.Println("The key exists")
   } else{
      fmt.Println("The key doesn't exist")
   }
}