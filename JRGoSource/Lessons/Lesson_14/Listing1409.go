package main

import "fmt"

type account struct {
   number string
   balance float64
}

// function defined with value receiver
func withdraw(acct account, value float64) account {
   if acct.balance >= value{
      acct.balance = acct.balance - value
   }
   return acct
}

func main() {
   acct := account{}
   acct.number="C21345345345355"
   acct.balance=159

   ptra :=&acct // create a pointer

   fmt.Println("Before: ", ptra)

   // This is not ok because the function accepts only value arguments
   // The withdraw statement won't execute and will throw an error
   withdraw(ptra,150)
 
   fmt.Println("After:" , ptra)
}