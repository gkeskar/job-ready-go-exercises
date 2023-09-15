package main

import "fmt"

func main() {
	//	display A to Z using for loop
	//	 array that stores chars A to Z
	//	use for loop and display array elements
	var arrayOfLetters [26]rune

	for i, _ := range arrayOfLetters {
		fmt.Print(string(rune('A'+i)), " ")
	}

	//for i := 0; i < 26; i++ {
	//	arrayOfLetters[i] = rune('A' + i)
	//	letter := string(arrayOfLetters[i])
	//	fmt.Print(letter, " ")
	//}

}
