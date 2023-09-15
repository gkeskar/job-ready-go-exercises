package main

import (
	"fmt"
)

func main() {
	var count int32
	fmt.Print("Enter the value:")
	fmt.Scan(&count)
	for i := 0; count > 0; i++ {
		if i == 0 {
			fmt.Println(i)

		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
			count = count - 1

		} else if i%3 == 0 {
			fmt.Println("fizz")

			count = count - 1

		} else if i%5 == 0 {
			fmt.Println("buzz")
			count = count - 1

		} else {
			fmt.Println(i)

		}
	}
	fmt.Println("TRADITION!!")
}
