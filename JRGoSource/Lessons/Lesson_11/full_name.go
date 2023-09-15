package main

import "fmt"

type name struct {
	firstName  string
	lastName   string
	middleName string
}

func GetFullName(inputName name) string {
	if inputName.middleName != "" {
		return inputName.firstName + inputName.middleName + "." + inputName.lastName
	} else {
		return inputName.firstName + " " + inputName.lastName
	}
}

func main() {
	var n = name{firstName: "John", lastName: "Doe"}

	fmt.Println(GetFullName(n))
}
