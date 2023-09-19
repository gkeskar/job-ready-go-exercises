package main

import "fmt"

type account struct {
   number string
   balance float64
}

// Function defined with pointer parameter
func withdraw(acct *account, value float64) account {
   if acct.balance >= value{
      acct.balance = acct.balance - value
   }
   return *acct
}

func main() {
   acct := account{}
   acct.number="C21345345345355"
   acct.balance=159
   ptra :=&acct // create a pointer
   fmt.Println(&ptra)

   fmt.Println("Before:", acct)
   // This statement will not execute and will throw an error
   withdraw(acct,150)
   fmt.Println("After:", acct)
}