package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int) 
   fmt.Println("Map value:", m)

   // add key-value pairs to the map
   m["Hi"] = 20
   m["How"] = 245
   fmt.Println("Updated map:", m)

   // Retrieve a value based on its key
   var num int = m["Hi"]
   fmt.Println("Value of num:", num)
}