package main

import "fmt"

func main() {
   fmt.Print("Enter your first name: ") //Print function displays
                                        //output in same line
   var firstName string
   fmt.Scanln(&firstName) // take input from user

   fmt.Print("Enter your last name: ")
   var lastName string
   fmt.Scanln(&lastName)

   fmt.Print("Your first name is: ")
   fmt.Println(firstName)

   fmt.Print("Your last name is: ")
   fmt.Println(lastName)
}