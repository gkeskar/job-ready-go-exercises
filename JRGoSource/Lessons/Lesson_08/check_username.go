package main

import "fmt"

var correctUsername, correctPasswd string

func checkUsernamePasswd(username string, password string) bool {
	if username == correctUsername && password == correctPasswd {
		return true
	} else {
		return false
	}
}

func main() {
	var inputUsername, inputPasswd string
	var isMatched bool

	correctUsername = "gkeskar"
	correctPasswd = "test123"

	fmt.Print("Enter Usernamee: ")
	fmt.Scan(&inputUsername)
	fmt.Print("Enter password: ")
	fmt.Scan(&inputPasswd)
	isMatched = checkUsernamePasswd(inputUsername, inputPasswd)
	fmt.Println("correct username, passwd entered?", isMatched)
}
