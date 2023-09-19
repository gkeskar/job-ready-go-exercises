package main

import "fmt"

type account struct {
   number string
   balance float64
}

// method with value receiver
func(acct account) withdraw(value float64) bool{
   if acct.balance >= value{
      acct.balance = acct.balance - value
      return true
   }
   return false
}

func main() {
   acct := account{}
   acct.number = "C21345345345355"
   acct.balance = 159

   fmt.Println(acct)
   fmt.Println(acct.withdraw(150)) // call the method defined on account
   fmt.Println(acct)
}