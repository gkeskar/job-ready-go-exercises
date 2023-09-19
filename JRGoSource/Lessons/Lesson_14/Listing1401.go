package main

import "fmt"

type Account struct {
   Number string
   Balance float64
}

// HaveEnoughBalance is a method defined on the struct, Account. 
// The receiver argument is (acct Account) which is separate from 
// the input argument list (value)
func (acct Account) HaveEnoughBalance(value float64) bool {
   return acct.Balance >= value
}

// HaveEnoughBalance2 is a simple function
func HaveEnoughBalance2(acct Account, value float64) bool {
   return acct.Balance >= value
}

func main() {
   a := Account{Number: "C21345345345355"}

   // call the method defined on Account
   fmt.Println("Method result:", a.HaveEnoughBalance(150))

   // call the function
   fmt.Println("Function result:", HaveEnoughBalance2(a, 150))
}