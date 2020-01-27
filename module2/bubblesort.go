package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(a_slice []int, index int) {
	a_slice[index], a_slice[index+1] = a_slice[index+1], a_slice[index]
}

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-2; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func main() {
	fmt.Println("Please input numbers separated with space:")
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	line_array := strings.Split(string(line), " ")
	var values []int
	for _, str := range line_array {
		num, _ := strconv.Atoi(str)
		values = append(values, num)
	}
	BubbleSort(values)
	fmt.Println(values)
}
