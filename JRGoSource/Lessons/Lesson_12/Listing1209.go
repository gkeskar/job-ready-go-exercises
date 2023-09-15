package main

import (
   "fmt"
)

func main() {
   var s []int   // create an empty slice 
   fmt.Println(s)
   s[0]=10       // this won't execute 
}
