package main

import(
   "fmt"
   "strings"
)

func DisplayUpper(x string) {
   fmt.Println("Original text:", x)
   fmt.Println("Revised text:", strings.ToUpper(x))
}

func main() {
   a := "elizabeth"

   DisplayUpper(a)
}