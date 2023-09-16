package main

import "fmt"

func main() {

	var mymap1 = make(map[string]bool)
	mymap1["day1"] = true
	mymap1["day2"] = true
	mymap1["day3"] = true
	mymap1["day4"] = true

	myMap := map[string]int{
		"day1":  12,
		"day2":  24,
		"day3":  30,
		"day5":  12,
		"day6":  3,
		"day7":  7,
		"day8":  15,
		"day9":  61,
		"day10": 98,
	}

	fmt.Println(mymap1)
	for key, value := range myMap {
		fmt.Print("key => ", key)
		fmt.Println(" value => ", value)
	}
}
