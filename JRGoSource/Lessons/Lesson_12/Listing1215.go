package main

import (
	"fmt"
)

func main() {
	var mySlice = new([10]int)[0:5]
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3
	mySlice[3] = 4
	mySlice[4] = 5

	fmt.Println(mySlice)
	newSlice := RemoveIndex(mySlice, 2)
	fmt.Println(newSlice)
}

func RemoveIndex(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
