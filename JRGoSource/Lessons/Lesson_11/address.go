package main

import (
	"fmt"
)

func main() {
	type address struct {
		firstName    string
		lastName     string
		addressLine1 string
		addressLine2 string
		city         string
		state        string
		country      string
		zipCode      int
	}

	var a = address{firstName: "John", lastName: "Doe", addressLine1: "123 My Street", city: "Chicago", country: "IL", zipCode: 12345}
	//fmt.Println(a)
	fmt.Println(a.firstName, a.lastName)
	fmt.Println(a.addressLine1)
	if a.addressLine2 != "" {
		fmt.Println(a.addressLine1)
	}
	fmt.Print(a.city, ",", a.state, " ", a.country, " ", a.zipCode)

}
