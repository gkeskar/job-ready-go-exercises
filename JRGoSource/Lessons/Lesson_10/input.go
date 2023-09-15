package main

import "fmt"

func main() {
	var nameptr *string
	var ageptr *string
	var genderptr *string

	fmt.Print("Enter your name: ")

	var name string
	fmt.Scanln(&name)
	nameptr = &name

	fmt.Print("Enter your age: ")
	var age string
	fmt.Scanln(&age)
	ageptr = &age

	fmt.Print("Enter your gender: ")
	var gender string
	fmt.Scanln(&gender)
	genderptr = &gender

	fmt.Print("Your first name is: ")
	fmt.Println(*nameptr)

	fmt.Print("Your age is: ")
	fmt.Println(*ageptr)

	fmt.Print("Your gender is: ")
	fmt.Println(*genderptr)
}
