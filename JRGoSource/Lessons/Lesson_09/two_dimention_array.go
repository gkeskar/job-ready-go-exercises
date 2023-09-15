package main

import "fmt"

func main() {
	matrix := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println(matrix)
	fmt.Println("length of main array:", len(matrix))

	fmt.Println(matrix[1])
	fmt.Println("length of second nested array:", len(matrix[1]))

	fmt.Println("last value in last nested array", matrix[3][3])
}
