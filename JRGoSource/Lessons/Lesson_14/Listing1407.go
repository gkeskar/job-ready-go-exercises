package main

import "fmt"

type SavingsAccount struct {
   number string
   balance float64
}

type CheckingAccount struct {
   number string
   balance float64
}

func(acct *SavingsAccount) Withdraw(value float64) bool {
   if acct.balance >= value{
      acct.balance = acct.balance - value
      return true
   }
   return false
}

func(acct *CheckingAccount) Withdraw(value float64) bool {
   if acct.balance >= value{
      acct.balance = acct.balance - value
      return true
   }
   return false
}

func main() {
   acct := SavingsAccount{}
   acct.number="S21345345345355"
   acct.balance=159

   fmt.Println("savings balance", acct)
   fmt.Println("withdraw from savings:", acct.Withdraw(150))
   fmt.Println("new savings balance", acct)

   acct2 := CheckingAccount{}
   acct2.number="C218678678345345355"
   acct2.balance=2000

   fmt.Println("checking balance", acct2)
   fmt.Println("withdraw from checking:", acct2.Withdraw(150))
   fmt.Println("new checking balance", acct2)
}