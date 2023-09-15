package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A function should take string and return the number of vowels
func countVowels(inputString string) int {
	stringOfVowels := strings.ToLower("AEIOU")
	count := 0
	for _, letter := range strings.ToLower(inputString) {
		fmt.Println(string(letter))
		if strings.Contains(stringOfVowels, string(letter)) {
			count += 1
		}
	}
	return count
}

// A function should take string and return the number of consonants
func countConsonants(inputString string) int {
	stringOfVowels := strings.ToLower("AEIOU")
	count := 0
	for _, letter := range strings.ToLower(inputString) {
		fmt.Println(string(letter))
		if !strings.Contains(stringOfVowels, string(letter)) {
			count += 1
		}
	}
	return count
}

func main() {

	// user promts
	fmt.Print("Enter a string: ")
	var inputString string
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')

	fmt.Println("Number of vowels in the string are:", countVowels(inputString))
	fmt.Println("Number of consonants in the string are:", countConsonants(inputString))
}
