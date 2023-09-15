package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isPime(number int) bool {
	if number == 2 {
		return true
	} else {

		limit := math.Sqrt(float64(number))
		for i := 2; i <= int(limit); i++ {
			if number%i == 0 {
				return false
			}
		}
	}
	return true
}

func factorial(number int) int64 {
	var output int64 = 1
	for i := number; i > 0; i-- {
		output = int64(output * int64(number))
		number = number - 1
	}
	return output
}

func main() {
	//	prompts user for an integer
	var num string
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	// remove leading zeros
	num = strings.TrimLeft(num, "0")

	//	display num of digits
	// convert into string and get length
	fmt.Println("The number of digits in the value entered:", len(string(num)))
	//	 first and last digit of num
	fmt.Println("The first, and the last digit of the number ", string(num[0]), string(num[len(num)-1]))
	var sum int
	var product int = 1
	for _, n := range num {
		numDigit, _ := strconv.Atoi(string(n))
		sum = sum + numDigit
		product = product * numDigit
	}

	fmt.Println("the sum of the digits in the number", sum)
	fmt.Println("the product of the digits in the number", product)
	numValue, _ := strconv.Atoi(num)
	if isPime(numValue) {
		fmt.Println("The number is prime")
	} else {
		fmt.Println("The number is not prime")
	}
	fmt.Println("The factorial of number is", factorial(numValue))
}
