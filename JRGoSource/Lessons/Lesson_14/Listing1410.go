package main

import "fmt"

type account struct {
   number string
   balance float64
}

// method defined with a pointer receiver
func(acct *account) withdraw(value float64) bool {
   if acct.balance>=value{
      acct.balance=acct.balance-value
      return true
   }
   return false
}

func main() {
   acct := account{number: "C21345345345355", balance: 159.0}

   fmt.Println("Before:", acct) 

   // The acct.withdraw will be interpreted by compiler as 
   // (&acct).withdraw(). This is ok because the method can 
   // accept both value and pointer receivers

   acct.withdraw(100)

   fmt.Println("After:", acct) 
}