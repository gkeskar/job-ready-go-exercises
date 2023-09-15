package main

import "fmt"

func main() {
   var username string = "chris";
   var password string = "dsxscg34"
   if (username == "mary" && password == "dsxscg34"){
      fmt.Println("This person has the right credentials.") 
   }
   if(username != "mary" || password != "dsxscg34"){
      fmt.Println("This person does not have the right credentials.")
   }
}