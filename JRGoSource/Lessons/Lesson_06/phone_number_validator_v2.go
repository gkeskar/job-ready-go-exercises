package main

import (
	"fmt"
	"unicode"
)

func hasAllDigits(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func is_valid_code(input_code string) bool {
	if string(input_code[0]) == "0" || string(input_code[0]) == "1" {
		return false
	}
	return true
}

func main() {

	var input_phone_number string
	var country_code bool
	var area_code, central_office_code, subscriber_number string
	fmt.Print("Enter phone number:")
	fmt.Scan(&input_phone_number)
	if !hasAllDigits(input_phone_number) {
		fmt.Println("Not a valid num as number has character")
		return
	}

	switch {
	case len(input_phone_number) < 10:
		fmt.Println("Not a valid num as num is less than 10 digits")
		return
	case len(input_phone_number) > 11:
		fmt.Println("Not a valid num as num is greater than 11 digits")
		return
	case len(input_phone_number) == 11:
		if string(input_phone_number[0]) != "1" {
			fmt.Println("Not a valid num as area code is not US area code starting with 1")
			return
		} else {
			fmt.Println("Valid phone num with US code")
			country_code = true
			input_phone_number = string(input_phone_number[1:])
		}
	case len(input_phone_number) == 10:
		country_code = false
	}

	area_code = input_phone_number[:3]            // first 3 digits
	central_office_code = input_phone_number[3:6] // next 3 digits
	subscriber_number = input_phone_number[6:]
	if !is_valid_code(area_code) || !is_valid_code(central_office_code) {
		fmt.Println("Not a valid num as area code or central office code can't start with 0 or 1")
		return
	}
	input_phone_number = area_code + "-" + central_office_code + "-" + subscriber_number
	if !country_code {
		fmt.Println("Valid number:", input_phone_number)
	} else {
		input_phone_number = "1" + "-" + input_phone_number
		fmt.Println("Valid number:", input_phone_number)
	}

}
