package main

import (
   "fmt"
   "math/rand"
)

func passGenerator() func() string {
   length := 10
   return func() string {
      pwd := ""
      for i := 0; i < length; i++ {
         // generate a number between 0 and 255 and convert it 
         // into its equivalent in UTF-8 
         randomChar := string(rand.Intn(256))
         pwd += randomChar
      }
      return pwd
   }
}

func main() {
   passGen := passGenerator()
   fmt.Println(passGen())
   fmt.Println(passGen())
   fmt.Println(passGen())
}