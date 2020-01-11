package main

import "fmt"

func main() {
    var num float64
    fmt.Println("Insert a floating point number.")
    _, err := fmt.Scan(&num)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(int(num))
}
