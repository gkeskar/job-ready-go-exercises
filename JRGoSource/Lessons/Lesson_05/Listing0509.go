package main

import "fmt"

func main() {
   var color string = "Yellow"

   switch(color){
   case "Red", "Blue", "Yellow":
      fmt.Println(color, "is a primary color")
   case "Orange", "Green", "Violet":
      fmt.Println(color, "is a secondary color")
   default:
      fmt.Println(color, "is not a primary or secondary color.")
   }
}