package main

import "fmt"

func main() {
	var input string
	fmt.Println("What season it it? ")
	fmt.Scan(&input)

	switch (input) {
	case "fall", "autumn":
		fmt.Println("I bet leaves are pretty there!")
	case "winter":
		fmt.Print("I hope you are ready for snow")
	case "spring":
		fmt.Println("I can smell the flowers")
	case "summer":
		fmt.Println("Make sure your AC is working")
	default:
		fmt.Println("I don't recongnize that season")
	}

}
