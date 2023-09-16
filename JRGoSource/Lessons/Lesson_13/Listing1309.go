package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int) 

   m["Hi"] = 20
   m["How"] = 245
   fmt.Println("Updated map:", m)

   // Retrieve a value based on its key
   var num int = m["BadKey"]
   fmt.Println("Value of num:", num)
}