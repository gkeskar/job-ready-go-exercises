package main

import "fmt"

func main() {
	//	computes length of string
	var inputString string
	var length int = 0
	fmt.Println("Enter String:")
	fmt.Scan(&inputString)
	//	input "abcde"
	// output 5
	for _, _ = range inputString {
		length = length + 1
	}
	fmt.Println(length)
}
