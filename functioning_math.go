package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var number int64
	var sqRoot float64
	var binary string

	fmt.Println("Enter the number: ")
	fmt.Scan(&number)
	sqRoot = math.Sqrt(float64(number))
	fmt.Println("You selected: ", number)
	fmt.Printf("The square root of your number is %.2f\n", sqRoot)
	binary = strconv.FormatInt(number, 2)
	fmt.Printf("The binary of your number is %s", binary)

}
