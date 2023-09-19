package main

import "fmt"

// Use the type keyword to create a new type, "Text", based 
// on the string type
type Text string

// This method has a receiver of type Text, as defined above
func (t Text) IsEmpty() bool {
   return len(t) <= 0
}

func main() {
   text := Text("Hi")
   fmt.Println("type value:", text)
   fmt.Println("method value:", text.IsEmpty())
}