package main

import "fmt"

func main() {
	var a, b int32
	a = 0
	b = 1
	fmt.Println("a < b", a < b)
	a = 3
	b = 1
	fmt.Println("a < b", a < b)
	a = 1
	b = 1
	fmt.Println("a < b", a < b)
	a = 1
	b = -2
	fmt.Println("a < b", a < b)
}
