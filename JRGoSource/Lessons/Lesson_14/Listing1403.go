package main

import "fmt"

// method attemps to use the string variable as the receiver
func (t string) IsEmpty() bool {
   return len(t) <= 0
}

func main() {
   text := s("Hi")
   fmt.Println(text)
   fmt.Println(text.IsEmpty())
}