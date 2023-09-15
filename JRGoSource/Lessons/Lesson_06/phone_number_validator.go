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

func main() {
	var input_phone_number string
	var area_code, central_office_code, subscriber_number string
	var country_code bool
	fmt.Print("Enter input phone number: ")
	fmt.Scan(&input_phone_number)

	// check if string has any chars or check if all digits
	if !hasAllDigits(input_phone_number) {
		fmt.Println("Not a valid num as number has character")
	} else {
		// check num less than 10
		if len(input_phone_number) < 10 {
			fmt.Println("Not a valid num as digits are less than 10")
			return
		}
		//	 check the length of string should be 10 digits or 11 digits with first digit as "1"
		if len(input_phone_number) > 11 {
			fmt.Println("Not a valid num as digits are greater than 11")
			return
		}
		if len(input_phone_number) == 11 {
			//	check for first digit sould be 1
			if string(input_phone_number[0]) != "1" {
				fmt.Println("Not a valid num as area code is not US area code starting with 1")
				return
			} else {
				fmt.Println("Valid phone num with US code")
				country_code = true
				input_phone_number = string(input_phone_number[1:])
			}
		}
		fmt.Println("valid phone number", input_phone_number)
		//	number has either 10 digits  or 11 digits starting with country code 1
		//	divide 10 digits into 3 digits of area code, 3 digits of central office code and 4 digits of subscriber num
		area_code = input_phone_number[:3]            // first 3 digits
		central_office_code = input_phone_number[3:6] // next 3 digits
		subscriber_number = input_phone_number[6:]
		fmt.Println("area_code: ", area_code)
		fmt.Println("central_office_code: ", central_office_code)
		fmt.Println("subscriber_number: ", subscriber_number)
		//	0 and 1 not allowed as first digit in area code and central office code
		if string(area_code[0]) == "0" || string(area_code[0]) == "1" {
			fmt.Println("Not a valid num as area code can't start with 0 or 1")
			return
		}
		if string(central_office_code[0]) == "0" || string(central_office_code[0]) == "1" {
			fmt.Println("Not a valid num as central_office_code can't start with 0 or 1")
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
}
