package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func BubbleSort(arr []int) []int {
	// Bubble sort implementation for an array in go
	n := len(arr)
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-2; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func StringToIntArray(arr_string, sep string) []int {
	/*
	   Takes a comma separated string of integers and returns and array of integers
	*/
	line_array := strings.Split(string(arr_string), sep)
	var values []int
	for _, str := range line_array {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		values = append(values, num)
	}
	return values
}

func OneDToTwoD(arr []int, splits int) [][]int {
	/*
	   Converts a 1D array to a 2D array for further sorting using go routines.
	   The splits define the number of smaller arrays the 1D array should be divided in. If the length of arr is odd, then the last array in the 2D array will be odd,
	   as the remaining element will be added to it.
	*/
	twoD := make([][]int, splits)
	size := len(arr) / splits
	for i := 0; i < splits; i++ {
		twoD[i] = arr[i*size : (i+1)*size]
		// Handling array of odd length, adding the remainder element to the last 2D array
		if (len(arr)%2 == 1) && i == splits-1 {
			twoD[i] = append(twoD[i], arr[len(arr)-1])
		}
	}
	return twoD
}

func subSort(arr []int, output chan int, wg *sync.WaitGroup) {
	// Sorts a smaller array and sends the sorted values to the output channel
	defer wg.Done() // 3
	sortedArr := BubbleSort(arr)
	for _, val := range sortedArr {
		output <- val
	}
}

func main() {
	// Take user input
	fmt.Println("Enter sequence of numbers (comma-separated)")
	reader := bufio.NewReader(os.Stdin)
	// Read the string
	line, _, _ := reader.ReadLine()
	// Convert string to an integer array
	num_array := StringToIntArray(string(line), ",")
	// Define number of splits for 1D to 2D conversion
	const max_n int = 4
	twoDArray := OneDToTwoD(num_array, max_n)

	var wg sync.WaitGroup // 1

	// Create a channel to store the sorted values that will be returned by multiple go routines
	sortedArrays := make(chan int, len(num_array))

	// Loop over the 2D array and create go routines for sorting the individual arrays inside the 2D array
	for _, arr := range twoDArray {
		wg.Add(1) // 2
		go subSort(arr, sortedArrays, &wg)
	}

	// Termination of the channel and waitGroup
	go func() {
		defer close(sortedArrays)
		wg.Wait()
	}()

	// Put the sorted values from the channel into finalArray and then sort finalArray
	var finalArray []int
	for arr := range sortedArrays {
		finalArray = append(finalArray, arr)
	}
	fmt.Println(BubbleSort(finalArray))
}
