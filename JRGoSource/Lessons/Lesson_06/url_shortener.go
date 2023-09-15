package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func random_two_digits() string {
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	randomNumber1 := generator.Intn(10)
	fmt.Println(randomNumber1)
	randomNumber2 := generator.Intn(10)
	fmt.Println(randomNumber2)
	twoRandonDigits := strconv.Itoa(randomNumber1) + strconv.Itoa(randomNumber2)
	fmt.Println(twoRandonDigits)
	return twoRandonDigits
}
func main() {
	var url string
	fmt.Print("Enter Url:")
	fmt.Scan(&url)

	//	imaginary domain main
	var domainName string = "surl.com"
	// pageIndentifier is first and fourth char of page plus 2 random number digits
	var pageIdentifier string
	parts := strings.Split(url, "/")
	pageName := parts[len(parts)-1]
	pageIdentifier = string(pageName[0]) + string(pageName[3]) + random_two_digits()
	shortUrl := "http://" + domainName + "/" + pageIdentifier
	fmt.Println(shortUrl)
}
