package main

import "fmt"

func main() {
	var taxableIncome float64 = 35000
	var taxDue float64

	if taxableIncome <= 9875 {
		taxDue = taxableIncome * 0.1
	} else {
		taxDue = (9875 * .1) + ((taxableIncome - 9875) * .12)
	}

	fmt.Print("Your tax due is: ")
	fmt.Println(taxDue)
}
