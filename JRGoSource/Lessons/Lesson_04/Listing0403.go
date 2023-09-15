package main

import "fmt"

func main() {
   var a int = 10
   var b int32 = 20
   var c byte = 15
   var d float32 = 0.05

   fmt.Println( int32(a) + int32(b) )     // int & int32
   fmt.Println( int32(b) + int32(c) )     // int32 & byte
   fmt.Println( int(a) + int(c) )         // int & byte
   fmt.Println( float32(a) + float32(d) ) // int & float32
   fmt.Println( float32(b) * float32(d) ) // int32 & float32
   fmt.Println( float32(c) / float32(d) ) // byte & float32
}