package main

import "fmt"

func main() {
   var score int = 88
   var grade string

   switch {
     case score > 90 :
        grade = "A"
     case ( score > 80 ) && ( score <= 90 ) :
        grade = "B"
     case ( score > 70 ) && ( score <= 80 ) :
        grade = "C"
     case ( score > 60 ) && ( score <= 70 ) :
        grade = "D"
     case score <= 80 :
        grade = "F"
     default:
        grade = "unknown"
   }

   fmt.Println("Your grade is: ", grade )
}