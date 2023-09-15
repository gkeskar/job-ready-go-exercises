package main

import "fmt"

func main() {
   var color string = "Blue"

   if (color == "Blue" ){
      fmt.Println("Blue like the sky")
   } else if (color == "Red") {
      fmt.Println("Red like the sun")
   } else if (color == "Green"){
      fmt.Println("Green like the trees")
   } else {
      fmt.Println("Please choose a valid color.")
   }
}