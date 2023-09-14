package main

import "fmt"

func main() {
	var interest, principal, rate float32
	var days int64

	fmt.Println("Enter principal ammount:")
	fmt.Scan(&principal)
	fmt.Println("Enter rate of interest: ")
	fmt.Scan(&rate)
	fmt.Println("Enter Num of days:")
	fmt.Scan(&days)

	interest = principal * rate * float32(days) / 365
	fmt.Println("Simple interest is: ", interest)

}
