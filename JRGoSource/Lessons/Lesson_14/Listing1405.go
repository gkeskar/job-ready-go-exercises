package main

import "fmt"

type account struct {
   number string
   balance float64
}

func(acct account) withdraw(value float64) account {
   if acct.balance >= value {
      acct.balance = acct.balance-value
   }
   return acct
}

func main() {
   acct := account{}
   acct.number = "C21345345345355"
   acct.balance = 159

   // Show initial values
   fmt.Println(acct)
   // assign the result of withdraw to acct
   acct = acct.withdraw(150)

   // The changes are now properly recorded in the caller acct
   fmt.Println(acct)
}