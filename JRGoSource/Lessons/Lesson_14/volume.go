package main

import (
	"fmt"
	"math"
)

// calculate the volume of the solid

type Cube struct {
	length float64
}

type Box struct {
	length float64
	width  float64
	height float64
}

type Sphere struct {
	radius float64
}

// Cone struct
type Cone struct {
	radius, height float64
}

// Pyramid struct
type Pyramid struct {
	baseArea, height float64
}

func (cube Cube) volume() float64 {
	return math.Pow(cube.length, 3)
}

func (box Box) volume() float64 {
	return box.length * box.width * box.height
}

func (sphere Sphere) volume() float64 {
	return 4.0 / 3.0 * math.Pi * math.Pow(sphere.radius, 3)
}

// Volume method for Cone
func (c Cone) volume() float64 {
	return (1.0 / 3.0) * math.Pi * math.Pow(c.radius, 2) * c.height
}

// Volume method for Pyramid
func (p Pyramid) volume() float64 {
	return 1.0 / 3.0 * p.baseArea * p.height
}

func main() {
	c := Cube{length: 12}
	fmt.Printf("Volume of cube: %.2f units³\n", c.volume())

	box1 := Box{
		length: 12,
		width:  3,
		height: 8,
	}
	fmt.Printf("Volume of box: %.2f units³\n", box1.volume())
	sphere1 := Sphere{
		radius: 2,
	}
	fmt.Printf("Volume of sphere: %.2f units³\n", sphere1.volume())

	cone := Cone{radius: 3, height: 5}
	fmt.Printf("Volume of cone: %.2f units³\n", cone.volume())

	pyramid := Pyramid{baseArea: 10, height: 7}
	fmt.Printf("Volume of pyramid: %.2f units³\n", pyramid.volume())
}
