package main

// we import the strconv package which allows us to parse data
// between different data types
import (
	"fmt"
	"strconv"
)

func main() {
	for a := 0; a < 10; a++ {
		if a%2 == 1 {
			// the Itoa function converts the int into its
			// equivalent (UTF-8) string
			fmt.Println(strconv.Itoa(a) + " is an odd number")
		}
	}
}
