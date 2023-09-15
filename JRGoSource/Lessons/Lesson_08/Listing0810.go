package main

import (
   "fmt"
   "math"
)

func main() {
   circleStuff := func(radius int) (d int, c float32) {
      d = radius * 2
      c = 2 * math.Pi * float32(radius)
      return
   }

   fmt.Println(circleStuff(5))
}