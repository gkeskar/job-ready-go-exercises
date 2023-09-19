package main

import "fmt"

type account struct {
   number string
   balance float64
}

// the method uses a pointer to account
func(acct *account) withdraw(value float64) bool{
   if acct.balance >= value{
      acct.balance = acct.balance-value
      return true
   }
   return false
}

func main() {
   acct := account{}
   acct.number="C21345345345355"
   acct.balance=159

   fmt.Println(acct)
   fmt.Println(acct.withdraw(150))
   fmt.Println(acct)
}