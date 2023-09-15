package main

import "fmt"

func main() {
	var numbers = [3]int{5, 6, 7}

	for i, vaule := range numbers {
		x := len(numbers) - 1 // 2
		numbers[x-i] = vaule  // 2, 1, 0
	}

	fmt.Println(numbers)
}
