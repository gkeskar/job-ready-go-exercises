package main

import (
	"fmt"
	"net/mail"
	"strings"
)

func main() {
	var input_email string

	fmt.Print("Enter email address:")
	fmt.Scan(&input_email)

	// use for loop to parse the string
	//   string should have only one @ and . after @, at least one char before @
	_, err := mail.ParseAddress(input_email)
	if err != nil {
		fmt.Printf("Error: %v is not a valid email address.\n", input_email)
	} else {
		fmt.Println("valid format: true")
		parts := strings.Split(input_email, "@")

		identifier := parts[0]
		domainPart := parts[1]

		domain := strings.Split(domainPart, ".")
		subdomain := domain[0]

		fmt.Println("Domain:", subdomain)
		fmt.Println("Identifier:", identifier)
	}
}
