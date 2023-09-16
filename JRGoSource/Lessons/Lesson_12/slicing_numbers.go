package main

import (
	"fmt"
	"math/rand"
	"time"
)

// function that creates a slice with a collection on random numbers 0 and 100
func createRandonNumberSlice() []int {
	rand.Seed(time.Now().UnixNano()) // To get a different set of numbers for each run.
	mySlice := make([]int, 10)
	// generate random numbers 0 to 100
	for i := range mySlice {
		mySlice[i] = rand.Intn(100)
	}
	return mySlice
}

// function to sort a slice in place(asceending or descending)
func sortSlice(mySlice []int, sortingDirection string) []int {
	// implement bubbleesort
	if sortingDirection == "descending" {
		for i := 0; i < len(mySlice)-1; i++ {
			for j := 0; j < len(mySlice)-1; j++ {
				if mySlice[j] < mySlice[j+1] {
					mySlice[j], mySlice[j+1] = mySlice[j+1], mySlice[j]
				}
			}
		}
	} else if sortingDirection == "ascending" {
		for i := 0; i < len(mySlice)-1; i++ {
			for j := 0; j < len(mySlice)-1; j++ {
				if mySlice[j] > mySlice[j+1] {
					mySlice[j], mySlice[j+1] = mySlice[j+1], mySlice[j]
				}
			}
		}
	}
	return mySlice
}

// function that takes as input 2 sorted slices and combine slices into a single sorted slice
func mergeSortedSlices(slice1 []int, slice2 []int, sortingDirection string) []int {
	mergedSortedSlice := make([]int, 0, len(slice1)+len(slice2)) // allocate memory for efficiency

	if sortingDirection == "descending" {
		i, j := 0, 0
		for i < len(slice1) && j < len(slice2) {
			if slice1[i] >= slice2[j] {
				mergedSortedSlice = append(mergedSortedSlice, slice1[i])
				i++
			} else {
				mergedSortedSlice = append(mergedSortedSlice, slice2[j])
				j++
			}
		}

		// Process remaining elements, if any.
		for i < len(slice1) {
			mergedSortedSlice = append(mergedSortedSlice, slice1[i])
			i++
		}

		for j < len(slice2) {
			mergedSortedSlice = append(mergedSortedSlice, slice2[j])
			j++
		}
	} else if sortingDirection == "ascending" {
		i, j := 0, 0
		for i < len(slice1) && j < len(slice2) {
			if slice1[i] <= slice2[j] {
				mergedSortedSlice = append(mergedSortedSlice, slice1[i])
				i++
			} else {
				mergedSortedSlice = append(mergedSortedSlice, slice2[j])
				j++
			}
		}
		// process remaining elements if any
		for i < len(slice1) {
			mergedSortedSlice = append(mergedSortedSlice, slice1[i])
			i++
		}
		for j < len(slice2) {
			mergedSortedSlice = append(mergedSortedSlice, slice2[j])
			j++
		}
	}

	return mergedSortedSlice
}

func main() {
	myslice := createRandonNumberSlice()
	fmt.Println("generated random slice: ", myslice)
	fmt.Println("sorting in descending order", sortSlice(myslice, "descending"))
	fmt.Println("sorting in ascending order", sortSlice(myslice, "ascending"))
	fmt.Println("___________")
	slice1 := createRandonNumberSlice()
	fmt.Print("slice1 ", slice1)
	sortedSlice1 := sortSlice(slice1, "descending")
	fmt.Println(" sorted in descending order: ", sortedSlice1)
	slice2 := createRandonNumberSlice()
	fmt.Print("slice2 ", slice2)
	sortedSlice2 := sortSlice(slice2, "descending")
	fmt.Println(" sorted in descending order: ", sortedSlice2)
	fmt.Println("Combine and sort slice1 and slice2 with descending", mergeSortedSlices(sortedSlice1, sortedSlice2, "descending"))
	fmt.Println("Combine and sort slice1 and sclie2 with ascending", mergeSortedSlices(sortedSlice1, sortedSlice2, "ascending"))
}
