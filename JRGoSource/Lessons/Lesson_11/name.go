package main

import "fmt"

type address struct {
	toName       name
	addressLine1 string
	addressLine2 string
	city         string
	state        string
	country      string
	zipCode      int
}

type name struct {
	firstName  string
	lastName   string
	middleName string
}

func main() {
	var n = name{firstName: "John", lastName: "Doe"}
	var a = address{toName: n, addressLine1: "123 My Street", city: "Chicago", country: "IL", zipCode: 12345}
	//fmt.Println(a)
	fmt.Println(a.toName.firstName, a.toName.lastName)
	fmt.Println(a.addressLine1)
	if a.addressLine2 != "" {
		fmt.Println(a.addressLine1)
	}
	fmt.Print(a.city, ",", a.state, " ", a.country, " ", a.zipCode)

	fmt.Println(n.firstName, n.lastName)
}
