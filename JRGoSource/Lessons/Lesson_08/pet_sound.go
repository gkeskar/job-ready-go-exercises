package main

import (
	"fmt"
	"strings"
)

var petSounds = map[string]string{
	"dog":   "woof",
	"cat":   "meow",
	"tiger": "roar",
	"mouse": "squeak",
	"cow":   "moo",
}

func petSound(animal string) (string, bool) {
	sound, ok := petSounds[strings.ToLower(animal)]
	return sound, ok
}

func main() {
	var pet string
	fmt.Print("Enter pet:")
	fmt.Scan(&pet)
	sound, ok := petSound(pet)
	if !ok {
		fmt.Println("pet is not recognized: ")
		return
	} else {
		fmt.Printf("A %s says %s\n", pet, sound)
	}
}
