package main

import "fmt"

func main() {
	numbers := []int{123, 111, 333, 777, 222, 999, 555, 888, 666, 444}

	// print each value in the array
	var ctr int
	for ctr = 0; ctr < len(numbers); ctr++ {
		fmt.Println(numbers[ctr])
	}

	fmt.Println("before:", numbers)

	var numbersptr [10]*int
	for ctr := 0; ctr < len(numbersptr); ctr++ {
		numbersptr[ctr] = &numbers[ctr]

	}

	for x := 0; x < len(numbersptr); x++ {
		for y := 0; y < len(numbersptr)-1; y++ {
			if *numbersptr[y+1] > *numbersptr[y] {
				temp := *numbersptr[y+1]
				*numbersptr[y+1] = *numbersptr[y]
				*numbersptr[y] = temp
			}
		}
	}

	// check thee next element if they if a[0] > a[1]
	// else move a[1] a[0]
	//  [123 111 333 777 222 999 555 888 666 444]
	// [999 888 777 555 666 444 333 222 222 123]
	// highesh value is a[0]
	// high = a[0]
	// compare if high all the elements

	fmt.Println("after:", numbers)

}
