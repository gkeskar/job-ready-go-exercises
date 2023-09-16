package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int) // define and initialize a map
   fmt.Println("Map value:", m)

   // add key-value pairs to the map
   m["Hi"] = 20
   m["How"] = 245

   fmt.Println("Updated map:", m)
}