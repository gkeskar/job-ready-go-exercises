package main

import "fmt"

func main() {
	var name string
	var age int
	var gender string

	var nameptr *string
	var genderptr *string
	var ageptr *int

	nameptr = &name
	genderptr = &gender
	ageptr = &age

	name = "milky"
	age = 4
	gender = "male"
	fmt.Println("print variabled directly:", name, age, gender)
	fmt.Println("Printing variables through pointers:", *nameptr, *ageptr, *genderptr)
}
