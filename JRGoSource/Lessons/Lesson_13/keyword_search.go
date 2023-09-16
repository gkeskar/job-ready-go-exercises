package main

import "fmt"

func main() {

	// map with 10 key value pair

	var produceCatageries = map[string]string{
		"apple":     "fruit",
		"onion":     "vegetable",
		"cheeze":    "dairy",
		"tomato":    "vegetable",
		"kiwi":      "fruit",
		"potato":    "vegetable",
		"straberry": "fruit",
		"banana":    "fruit",
		"spinach":   "vegenatble",
		"eggs":      "poultry",
	}

	fmt.Println(produceCatageries)
	// prompt user to input a search term
	var input string
	for input != "exit" {
		fmt.Println("Enter a search:")
		fmt.Scan(&input)
		// display all key value pairs that include serach either in key or value
		if input == "exit" {
			break
		} else {
			var found bool
			for produce, catagery := range produceCatageries {

				if input == produce || input == catagery {
					fmt.Println(input, "is found in key", produce, "with value", catagery)
					found = true
				}
			}
			if !found {
				fmt.Println(input, " is not found in any of key value pairs")
				fmt.Println("Enter another search or type exit to quit")
			}
		}
	}

}
