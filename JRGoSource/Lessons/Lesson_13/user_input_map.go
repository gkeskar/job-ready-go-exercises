package main

import (
	"fmt"
	"strconv"
)

func main() {

	var dayCollection = make(map[int]float64)

	var input string
	var inputKey int
	var inputValue float64

	for input != "quit" {
		fmt.Println("Enter the keys or quit to exit")
		fmt.Scan(&input)
		if input != "quit" {
			inputKey, _ = strconv.Atoi(input)
			_, ok := dayCollection[inputKey]
			if ok == true {
				fmt.Println("Key already exists")
				continue
			}
		} else if input == "quit" {
			break
		}

		fmt.Println("Enter the value or quit to exit:")
		fmt.Scan(&input)
		if input != "quit" {
			inputValue, _ = strconv.ParseFloat(input, 64)
			dayCollection[inputKey] = inputValue
		}

	}
	fmt.Println(dayCollection)
}
