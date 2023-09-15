package main

import "fmt"

func main() {
	var accountBalance int = 0

	if (accountBalance > 0) == false {
		fmt.Println("This person's bank account has no money.")
	} else {
		fmt.Println("Balance verification is complete.")
	}
}
