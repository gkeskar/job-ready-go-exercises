package main

import (
	"fmt"
	"math"
)

func main() {

	var numberFloat float64
	fmt.Print("Enter the number: ")
	fmt.Scan(&numberFloat)

	fmt.Println(math.Trunc(numberFloat))

}
