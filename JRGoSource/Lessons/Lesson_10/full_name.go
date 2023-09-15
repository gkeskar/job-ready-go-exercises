package main

import "fmt"

func full_name(fullNameptr *string, firstName string, lastName string) string {
	fullName := firstName + " " + lastName
	fullNameptr = &fullName
	return *fullNameptr
}
func main() {
	var fullName string
	var fullNameptr *string
	fullName = full_name(fullNameptr, "Milky", "Aradhye")
	fmt.Println(fullName)
}
