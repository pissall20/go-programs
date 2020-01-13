package main

import (
	"fmt"
	"sort"
	"strconv"
	"os"
	"strings"
)


func main() {
	slice := make([]int, 0, 3)

	for {

		var numString string

		fmt.Println("Enter a number:")
		fmt.Scan(&numString)
		
		numString =  strings.ToLower(numString)

		if numString == "x"{
			os.Exit(0)
		}

		num, err := strconv.Atoi(numString)

		if err != nil {
			fmt.Println("You have entered the wrong output")
			os.Exit(1)
		}

		slice = append(slice, num)

		sort.Ints(slice)
		fmt.Println(slice)
	}
}
