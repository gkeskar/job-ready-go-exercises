package main

import "fmt"

func main() {
   var number int = 4

   switch (number) {
   case 10 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 9 :
      fmt.Println("...", number , "...")
      number -= 1
      fallthrough
   case 8 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 7 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 6 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 5 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 4 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 3 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 2 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 1 :
      fmt.Println("...", number, "...")
      number -= 1
      fallthrough
   case 0 :
      fmt.Println("*** BOOM ***")
   default:
      fmt.Println("Try a number from 1 to 10!")
   }
}