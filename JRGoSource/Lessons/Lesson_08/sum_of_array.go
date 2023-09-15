package main

import "fmt"

// Sum of Array Elements: Write a program to find the sum of an array using

func sum(a []int, n int) int {
	if n <= 0 {
		return 0
	} else {
		return sum(a, n-1) + a[n-1]
	}

}

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(array, len(array)))
}
