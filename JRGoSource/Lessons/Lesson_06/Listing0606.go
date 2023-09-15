package main

import (
	"fmt"
	"strconv"
)

func main() {
	var power2 int64 = 1
	var a int64 = 0
	for a < 10 {
		fmt.Println("2 to the power of " + strconv.FormatInt(a, 10) +
			" is equal to " + strconv.FormatInt(power2, 10))
		power2 += power2
		a++
	}
}
