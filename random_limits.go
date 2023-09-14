package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var input int
	fmt.Println("Enter integer number: ")
	fmt.Scan(&input)

	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	fmt.Println(generator.Intn(input))
}
