package main

import (
	"fmt"
)


func Swap(a_slice []int, index int) {
	a_slice[index], a_slice[index+1] = a_slice[index + 1], a_slice[index]
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
	s1 := []int{1, 5, 2, 47, 7, 8, 0, 6, 21}

	BubbleSort(s1)
	fmt.Println(s1)
}
