package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fuction that takes as an input int, n, and returns an array of length n with randon integers
func generateNumbers(n int) []int {
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)

	var arryOfIntegers = make([]int, n)
	for i := 0; i < n; i++ {
		element := generator.Intn(200) - 100
		arryOfIntegers[i] = element
	}
	return arryOfIntegers
}

type IntegerArray []int

// function to compute the max of an array of int and its index
func (integers IntegerArray) computeMax() (int, int) {
	var max int
	var indexOfmax int
	for i, x := range integers {
		if x > max {
			max = x
			indexOfmax = i
		}
	}
	return max, indexOfmax
}

// function to compute the min of an array of int and its index
func (integers IntegerArray) computeMin() (int, int) {
	var min int
	var indexOfmin int
	min = integers[0]
	for i, x := range integers {
		if x <= min {
			min = x
			indexOfmin = i
		}
	}
	return min, indexOfmin
}

// sort array in descending order and return sorted array in new seperate array
func (integers IntegerArray) sortedArray() []int {
	var n = len(integers)
	var sortedIntegers = make([]int, n)
	copy(sortedIntegers, integers) // use copy function to avoid altering original integers slice
	// use bubble sort
	for i := 0; i < len(sortedIntegers)-1; i++ {
		for j := 0; j < len(sortedIntegers)-1; j++ {
			if sortedIntegers[j] < sortedIntegers[j+1] {
				sortedIntegers[j], sortedIntegers[j+1] = sortedIntegers[j+1], sortedIntegers[j]
			}
		}
	}
	return sortedIntegers
}

// sort array in ascending order and return sorted array in new seperate array
func (integers IntegerArray) sortedArrayAscending() []int {
	var n = len(integers)
	var sortedIntegers = make([]int, n)
	copy(sortedIntegers, integers) // use copy function to avoid altering original integers slice
	// use bubble sort
	for i := 0; i < len(sortedIntegers)-1; i++ {
		for j := 0; j < len(sortedIntegers)-1; j++ {
			if sortedIntegers[j] > sortedIntegers[j+1] {
				sortedIntegers[j], sortedIntegers[j+1] = sortedIntegers[j+1], sortedIntegers[j]
			}
		}
	}
	return sortedIntegers
}

// compute the mean of the integers
func (integers IntegerArray) computeMean() int {
	var sum int
	for _, x := range integers {
		sum = sum + x
	}
	mean := sum / len(integers)
	return mean
}

// compute median of an array
// 1, 2, 3, 4, 5   :5 / 2  = 2
// 1, 2, 3, 4, 5, 6   :6 / 2  = 3
func (integers IntegerArray) computeMedian() float64 {
	//First, simply sort the array. Then, check if the number of elements present in the array is even or odd. If odd, then simply return the mid value of the array. Else, the median is the average of the two middle values.
	var median float64
	sortedIntegers := integers.sortedArrayAscending()
	if len(sortedIntegers)%2 == 1 {
		median = float64(sortedIntegers[len(integers)/2])
	} else {
		median = float64(sortedIntegers[len(integers)/2]+sortedIntegers[len(integers)/2+1]) / 2.0

	}
	return median
}

// return positive numbers in the array
func (integers IntegerArray) positiveIntegers() []int {
	var positiveIntegers []int
	for _, x := range integers {
		if x > 0 {
			positiveIntegers = append(positiveIntegers, x)
		}
	}
	return positiveIntegers

}

// return negative numbers in the array
func (integers IntegerArray) negativeIntegers() []int {
	var negativeIntegers []int
	for _, x := range integers {
		if x < 0 {
			negativeIntegers = append(negativeIntegers, x)
		}
	}
	return negativeIntegers

}

// compute the longest sequence of sorted numbers in asceding or descending order  and return the new array
// ex input [1, 45, 67, 87, 6, 57, 0]
// ex input [ 45, 1, 67, 87, 6, 57, 0]
// output [1, 45, 67, 87]

func (integers IntegerArray) longestOrderedList() []int {
	var longestAscSequence, longestDescSequence []int
	var currentAscSequence, currentDescSequence []int

	for i := 0; i < len(integers); i++ {
		if i == 0 || integers[i-1] > integers[i] {
			currentAscSequence = integers[i : i+1] // start a new asc sequence
		} else {
			currentAscSequence = append(currentAscSequence, integers[i]) // add to current asc sequence
		}

		if len(currentAscSequence) > len(longestAscSequence) {
			longestAscSequence = currentAscSequence // new longest asc sequence found
		}

		if i == 0 || integers[i-1] < integers[i] {
			currentDescSequence = integers[i : i+1] // start a new desc sequence
		} else {
			currentDescSequence = append(currentDescSequence, integers[i]) // add to current desc sequence
		}

		if len(currentDescSequence) > len(longestDescSequence) {
			longestDescSequence = currentDescSequence // new longest desc sequence found
		}
	}

	if len(longestAscSequence) >= len(longestDescSequence) {
		return longestAscSequence // return longest Ascending Sequence
	} else {
		return longestDescSequence // return longest Descending Sequence
	}
}

// function to reemove duplicates and return the unique elements

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func (integers IntegerArray) removeDuplicates() []int {
	i := 0
	for i < len(integers) {
		for j := i + 1; j < len(integers); {
			if integers[i] == integers[j] {
				integers = RemoveIndex(integers, j)
			} else {
				j++
			}
		}
		i++
	}
	return integers
}

func main() {
	array := generateNumbers(15)
	numbers := IntegerArray(array)
	fmt.Println(array)
	max, index := numbers.computeMax()
	fmt.Println("max is", max, "and index of max is:", index)
	min, indexOfmin := numbers.computeMin()
	fmt.Println("min is", min, "and index of min is:", indexOfmin)
	fmt.Println("sorted in descending", numbers.sortedArray())
	fmt.Println("original array", array)
	fmt.Println("sorted in ascending", numbers.sortedArrayAscending())
	fmt.Println("original array", array)
	fmt.Println("mean of integers", numbers.computeMean())
	fmt.Println("median of integers", numbers.computeMedian())
	fmt.Println("positive integers", numbers.positiveIntegers())
	fmt.Println("negative integers", numbers.negativeIntegers())
	fmt.Println("longest sequence of integers", numbers.longestOrderedList())
	fmt.Println("removing duplicates", numbers.removeDuplicates())
}
