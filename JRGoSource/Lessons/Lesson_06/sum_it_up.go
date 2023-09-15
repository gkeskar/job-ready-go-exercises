package main

import "fmt"

func main() {
	//	 sum all numbers between 0 and 100
	//	 create an array with all numbers from 0 to 100 using a loop
	//	 add those number
	//	 create a variable with num and sum
	//	 start a loop
	//	 all next num to sum

	var sum int = 0
	for i := 0; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
