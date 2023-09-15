package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var size string
	var coffee_type string
	var flavor string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want small, medium, or large? ")
	fmt.Scanln(&size)

	fmt.Print("Do you want brewed, espresso, or cold brew? ")
	coffee_type, _ = reader.ReadString('\n')
	coffee_type = strings.Replace(coffee_type, "\n", "", -1)
	//fmt.Scanln(&coffee_type)

	fmt.Print("Do you want hazelnut, vanilla, caramel or chocolate? ")
	fmt.Scanln(&flavor)
	if flavor == "" {
		flavor = "no"
	}

	fmt.Printf("You asked a %s cup of %s coffe with %s syrup \n", size, coffee_type, flavor)

	var cost float32
	var totalCost float32

	switch size {
	case "small":
		cost = cost + 2.0
	case "medium":
		cost = cost + 3.0
	case "large":
		cost = cost + 4.0
	default:
		cost = 0.0
	}

	switch coffee_type {
	case "brewed":
		cost = cost + 0.0
	case "espresso":
		cost = cost + 0.50
	case "cold brew":
		cost = cost + 1.0
	default:
		cost = cost + 0.0
	}

	switch flavor {
	case "hazelnut", "caramel", "vanilla", "chocolate":
		cost = cost + 0.50
	case "None", "no":
		cost = cost + 0.0
	}
	totalCost = cost + cost*0.15
	fmt.Printf("Your cup of cofee costs %.2f \n", cost)
	fmt.Printf("The price with tip is %.2f \n", totalCost)
}
