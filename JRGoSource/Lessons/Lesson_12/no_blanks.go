package main

import (
	"fmt"
	"strings"
)

func removeValueFromSlice(mySlice []string, i int) []string {
	return append(mySlice[:i], mySlice[i+1:]...)
}

func main() {
	mySlice := make([]string, 11)

	mySlice[0] = "Happy"
	mySlice[1] = "Sneezy"
	mySlice[2] = "Grumpy"
	mySlice[3] = "Fred"
	mySlice[4] = "Doc"
	mySlice[5] = "Sleepy"
	mySlice[6] = "Dopey"
	mySlice[7] = "Bashful"
	mySlice[8] = "Bashful"
	mySlice[9] = ""
	mySlice[10] = ""

	fmt.Println("Before", mySlice)
	mySlice = removeValueFromSlice(mySlice, 3)
	fmt.Println("After removing Fred ", mySlice)

	for i := len(mySlice) - 1; i >= 0; i-- {

		if len(strings.TrimSpace(mySlice[i])) == 0 {
			fmt.Println("blank indexes are", i)
			mySlice = removeValueFromSlice(mySlice, i)
		}
	}

	fmt.Println("After removing blanks:", mySlice)
}
