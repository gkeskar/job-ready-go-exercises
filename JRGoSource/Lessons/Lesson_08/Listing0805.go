package main

import "fmt" 

func circleStuff(radius int) (int, float32) {
   d := radius * 2
   c := 2 * 3.14 * float32(radius)
   return d, c 
}

func main() {
   diameter, circumference := circleStuff(5)

   fmt.Println("diameter:", diameter)
   fmt.Println("circumference:", circumference)
}