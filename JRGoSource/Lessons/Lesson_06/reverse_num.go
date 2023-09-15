package main

import "fmt"

func main() {
	var inputNum string
	fmt.Print("Enter the num:")
	fmt.Scan(&inputNum)
	var outputNum string

	for i := len(inputNum) - 1; i >= 0; i-- {
		outputNum = outputNum + string(inputNum[i])
	}

	fmt.Println(outputNum)
}
