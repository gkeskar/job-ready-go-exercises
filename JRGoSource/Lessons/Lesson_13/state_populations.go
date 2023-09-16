package main

import "fmt"

func main() {

	var statePopulations = map[string]int64{
		"FL": 22244823,
		"CA": 39029342,
		"NY": 19677151,
		"TX": 30029572,
	}

	fmt.Println(statePopulations)
	var highestPopulationState string
	var highestPopulation int64
	var lowestPopulationState string
	var lowestPopulation int64
	lowestPopulation = statePopulations["FL"]
	lowestPopulationState = "FL"
	//iterate through map and determine state with higest population
	var sum int64
	for state, population := range statePopulations {
		if population > highestPopulation {
			highestPopulation = population
			highestPopulationState = state
		}

		if population < lowestPopulation {
			lowestPopulation = population
			lowestPopulationState = state
		}
		sum = sum + population
	}
	fmt.Println("State with the highest population is", highestPopulationState, "with a population of", highestPopulation)
	fmt.Println("State with the lowest population is", lowestPopulationState, "with a population of", lowestPopulation)
	fmt.Println("average population of state:", sum/int64(len(statePopulations)))
	//iterate through map and determine state with lowest population
}
