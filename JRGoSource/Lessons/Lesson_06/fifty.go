package main

import "fmt"

func main() {
	//	display all numbers divisible by 50 between 100 and 10000(inclusive)

	//for i := 50; i <= 10000; i++ {
	//	if i%50 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	//	using for and range
	// create array and then use range
	numbers := make([]int, 1+(10000-50))
	for i := range numbers {
		numbers[i] = 50 + i
		if numbers[i]%50 == 0 {
			fmt.Println(numbers[i])
		}
	}

}
