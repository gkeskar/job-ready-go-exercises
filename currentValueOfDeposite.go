package main

import (
	"fmt"
	"math"
)

func main() {
	var numerator, v, p, r float64

	var n, t int
	var power int

	fmt.Print("Enter Initial deposit: ")
	fmt.Scan(&p)
	fmt.Print("Enter Interest Rate: ")
	fmt.Scan(&r)
	fmt.Print("Enter number of times interest is calculated: ")
	fmt.Scan(&n)
	fmt.Print("Enter number of years since the initial deposit is calculated: ")
	fmt.Scan(&t)

	numerator = p * (1 + r/float64(n))
	power = n * t

	v = math.Pow(numerator, float64(power))
	fmt.Printf("On Initial Deposit %.2f$ with interest rate of %.2f percent, for %d years, %d number of times interest is calculated. Your value is %.4f ", p, r, n, t, v)

}
