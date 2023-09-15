package main

import "fmt"

func main() {

   var username string = "chris";
   var password string = "dsxscg34"

   if (username == "john" && password == "dsxscg34"){
      fmt.Println("This person has the right credentials.")
   } else {
      // the else must follow the closing bracket of the if statement
      // and NOT appear on a new line
      fmt.Println("This person does not have the right credentials.")
   }
}