package main

import "fmt"

func main() {
   var a, b int16 = 10, 200
   fmt.Println("a =",a)
   fmt.Println("b =",b)

   fmt.Println("a & b:", a & b)   // binary AND
   fmt.Println("a | b:", a | b)   // binary OR
   fmt.Println("a ^ b:", a ^ b)   // binary XOR
   fmt.Println("a << b:", a << b) // binary left shift
   fmt.Println("a >> b:", a >> b) // binary right shift
}