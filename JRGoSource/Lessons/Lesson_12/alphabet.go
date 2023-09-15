package main

import (
	"fmt"
	"strings"
)

func main() {
	var alphabet [26]string
	for i := 0; i < 26; i++ {
		alphabet[i] = string(rune('A' + i))
	}

	var sliceOfConsonants, sliceOfVowels []string
	sliceOfVowels = []string{"A", "E", "I", "O", "U"}

	for _, letter := range alphabet {
		//convert our slice into a single string ("AEIOU") that contains all vowel characters
		stringsOfVowels := strings.Join(sliceOfVowels, "")
		if !strings.Contains(letter, stringsOfVowels) {
			sliceOfConsonants = append(sliceOfConsonants, letter)
		}
	}
	fmt.Println(alphabet, sliceOfVowels, sliceOfConsonants)
}
