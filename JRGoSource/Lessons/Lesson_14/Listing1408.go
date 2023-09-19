package main

import "fmt"

type account struct {
   number string
   balance float64
}

// method defined with value receiver
func(acct account) withdraw(value float64) bool {
   if acct.balance >= value{
      acct.balance = acct.balance - value
      return true
   }
   return false
}

func main() {
   acct := account{number: "C21345345345355", balance: 159.0}

   ptra :=&acct // create a pointer

   fmt.Println("Before:", ptra)

   // This is ok because the method can accept both a value and 
   // a pointer receiver.
   ptra.withdraw(100)
 
   fmt.Println("After: ", ptra)
}