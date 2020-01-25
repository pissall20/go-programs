package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
    "bufio"
)


func Sort(arr []int) {
    n := len(arr)
    for i:= 0; i <= n-1; i++{
        for j:=0; j <= n-2; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func StringToIntArray(arr_string, sep string) []int {
	line_array := strings.Split(string(arr_string), sep)
	var values []int
	for _, str := range line_array {
		num, _ := strconv.Atoi(str)
		values = append(values, num)
	}
        return values
}

func OneDToTwoD(arr []int, splits int) [][]int {
    twoD := make([][]int, splits)
    size := len(arr) / splits
    for i:= 0; i < splits; i++{
        twoD[i] = arr[i*size : (i+1) * size]
    }
    return twoD
}

func main() {
    fmt.Println("Enter sequence of numbers (comma-separated)")
    reader := bufio.NewReader(os.Stdin)
    line, _, _ := reader.ReadLine()
    num_array := StringToIntArray(string(line), ",")
    twoDArray := OneDToTwoD(num_array, 4)
    fmt.Println(twoDArray)
}
