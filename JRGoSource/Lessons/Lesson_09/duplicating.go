package main

import "fmt"

func main() {
	var numbers = [...]int{7, 5, 10, 15, 30, 50}

	var numbersDuplicated [6]int

	for x, value := range numbers {
		if x%2 == 1 {
			value = numbers[x-1]
		}
		numbersDuplicated[x] = value

	}
	fmt.Println(numbers)
	fmt.Println(numbersDuplicated)
}
