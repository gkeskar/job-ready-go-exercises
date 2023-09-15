package main

import "fmt"

func main() {
	var input float64
	fmt.Println("How much money you have in your wallet? ")
	fmt.Scan(&input)

	if input >= 20 {
		fmt.Println("You are rich")
	} else {
		fmt.Println("You are broke")
	}
}
