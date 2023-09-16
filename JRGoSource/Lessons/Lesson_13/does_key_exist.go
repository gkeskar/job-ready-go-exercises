package main

import "fmt"

func main() {
	// create a map with atleast 5 key value pairs

	DaysOfTheWeek := map[int]string{
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Wednesday",
		5: "Thursday",
	}

	fmt.Println(DaysOfTheWeek)
	// prompts user to enter the key
	fmt.Print("Enter the key to add: ")
	var inputKey int
	fmt.Scan(&inputKey)
	fmt.Println("You entered key ", inputKey)
	// check if key alrealy exists
	_, ok := DaysOfTheWeek[inputKey]
	if ok == true {
		fmt.Println("Key already exists with value:", DaysOfTheWeek[inputKey])
		fmt.Println("Do you want to delete it? ")
		var deleteKey string
		fmt.Scan(&deleteKey)
		if deleteKey == "yes" {
			delete(DaysOfTheWeek, inputKey)
			fmt.Println("updated map:", DaysOfTheWeek)
		}

	} else {
		fmt.Println("Key is not used")
		fmt.Println("Enter value:")
		var inputValue string
		fmt.Scan(&inputValue)
		DaysOfTheWeek[inputKey] = inputValue
		fmt.Println("updated map:", DaysOfTheWeek)
	}
}
