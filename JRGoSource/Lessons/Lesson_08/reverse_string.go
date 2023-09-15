package main

// Write a program to reverse a string using recursion.

import "fmt"

func reverse(s string) string {
	if len(s) == 1 {
		return s
	} else {
		return string(s[len(s)-1]) + reverse(s[:len(s)-1])
	}

}

func main() {

	s := "abcdefg"
	fmt.Println(reverse(s))
}
