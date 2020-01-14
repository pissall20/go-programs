package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string
	fmt.Println("Please input a string:")
	_, err := fmt.Scan(&myString)
	if err != nil {
		fmt.Println(err)
	}
	if strings.HasPrefix(myString, "i") && strings.HasSuffix(myString, "n") && strings.Contains(myString, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
