package main

import (
   "fmt"
   "math/rand"
   "time"
)

func main() {
   ns := rand.NewSource(time.Now().UnixNano())
   generator := rand.New(ns)

   fmt.Println(generator.Intn(100))
   fmt.Println(generator.Intn(100))
}