package main

import (
   "fmt"
)

func main() {
   var m map[string]int

   fmt.Println(m) // (Reading) Displays the nil map as empty 
   m["Hi"] = 1 // (Writing) Throws an error
}