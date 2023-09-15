package main

import "fmt"

// The Fibonacci series is a sequence of numbers where the next number is found by adding up the two numbers before it. The sequence starts 0, 1, 1, 2, 3, 5, 8, 13,...

func main() {
	n := 10

	fibonaacchi := func(n int) int {
		if n <= 1 {
			return 1
		} else {
			return fibonaacchi(n-1) + fibonaacchi(n-2)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(fibonaacchi(i))
	}
}
