package main

import (
   "fmt"
)

func main() {
   FreqOccurrence := map[string]int{
      "hi": 23,
      "hello": 2,
      "hey": 4,
      "weather": 1.5,
      "greet": 35,
   }

   fmt.Println("Map value:", FreqOccurrence)
}