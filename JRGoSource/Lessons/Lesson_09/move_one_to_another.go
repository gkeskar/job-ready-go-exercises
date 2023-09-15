package main

import "fmt"

func main() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var reversedNumbers [10]int

	for i, value := range numbers {
		x := len(numbers)                // 3
		reversedNumbers[(x-1)-i] = value // 2, 1, 0

	}

	fmt.Println(reversedNumbers)
}
