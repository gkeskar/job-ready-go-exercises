package main

import "fmt"

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

	var AddressBook [4]address
	var a = address{firstName: "John", lastName: "Doe", addressLine1: "123 My Street", city: "Chicago", country: "IL", zipCode: 12345}
	var b = address{firstName: "Pat", lastName: "Smith", addressLine1: "123 My Street", addressLine2: "my boulevard", city: "Chicago", country: "IL", zipCode: 12345}
	var c = address{firstName: "Mark", lastName: "Anderson", addressLine1: "123 My Street", city: "Chicago", country: "IL", zipCode: 12345}
	var d = address{firstName: "Rachel", lastName: "Dean", addressLine1: "123 My Street", city: "Chicago", country: "IL", zipCode: 12345}
	AddressBook[0] = a
	AddressBook[1] = b
	AddressBook[2] = c
	AddressBook[3] = d

	for _, address := range AddressBook {
		fmt.Printf("%s %s\n", address.firstName, address.lastName)
		fmt.Printf("%s\n", address.addressLine1)
		if address.addressLine2 != "" {
			fmt.Printf("%s\n", address.addressLine2)
		}
		fmt.Printf("%s,%s %s %d\n", address.city, address.state, address.country, address.zipCode)
		fmt.Println("________________")

	}
}
