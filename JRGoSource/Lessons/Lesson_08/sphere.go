package main

import "fmt"

func sphereStuff(radius int) (int, float32, float32) {
	d := radius * 2
	c := 2 * 3.14 * float32(radius)
	a := 4 * 3.14 * float32(radius) * float32(radius)
	return d, c, a
}

func main() {
	diameter, circumference, area := sphereStuff(5)

	fmt.Println("diameter:", diameter)
	fmt.Println("circumference:", circumference)
	fmt.Println("area:", area)
}
