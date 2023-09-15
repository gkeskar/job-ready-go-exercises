package main

import (
   "fmt" 
   "math"
)

func circleStuff(radius int) (d int, c float32) {
   d = radius * 2
   c = 2 * math.Pi * float32(radius)
   return 
}

func main() {
   diameter, _ := circleStuff(5)
   fmt.Println("diameter:", diameter)
}