package main

import (
	"fmt"
	"strconv"
)

func main() {
	var width, length, height string
	var widthInt, lengthInt, heightInt int

	fmt.Print("Enter the width: ")
	fmt.Scan(&width)
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)
	fmt.Print("Enter the height: ")
	fmt.Scan(&height)

	// convert to int
	widthInt, _ = strconv.Atoi(width)
	lengthInt, _ = strconv.Atoi(length)
	heightInt, _ = strconv.Atoi(height)

	fmt.Println(widthInt * lengthInt * heightInt)
}