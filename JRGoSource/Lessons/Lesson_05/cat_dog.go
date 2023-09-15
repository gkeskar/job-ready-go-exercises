package main

import "fmt"

func main() {
	var having_cats string
	var having_dogs string

	fmt.Println("Do you have cats? ")
	fmt.Scan(&having_cats)

	fmt.Println("Do you have dogs? ")
	fmt.Scan(&having_dogs)

	if having_cats == "Yes" && having_dogs == "Yes" {
		fmt.Println("You must really love pets!")
	}
	if having_cats == "No" && having_dogs == "Yes" {
		fmt.Println("May be you need more pets.")
	}
	if having_cats == "No" && having_dogs == "No" {
		fmt.Println("May be you need more pets.")
	}
	if having_cats == "Yes" && having_dogs == "No" {
		fmt.Println("May be you need more pets.")
	}

}
