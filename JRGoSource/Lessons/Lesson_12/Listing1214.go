package main

import (
   "fmt"
   "reflect"
)

func main() {
   var mySlice = new ([10]string) [0:8]
   mySlice[0] = "Happy"
   mySlice[1] = "Sneezy"
   mySlice[2] = "Grumpy"
   mySlice[3] = "Bashful"
   mySlice[4] = "Doc"
   mySlice[5] = "Sleepy"
   mySlice[6] = "Dopey"
   mySlice[7] = "Fred"

   fmt.Println("slice type:", reflect.TypeOf(mySlice))
   fmt.Println(mySlice)
   fmt.Println(mySlice[2])
   fmt.Println(mySlice[2:5])
}