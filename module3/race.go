package main

import (
	"fmt"
)

func main() {
	// Define a counter for printing
	i := 0
	// Run an infinite loop to demonstrate a race condition as and when it happens
	for {
		var x, y int

		// Set x to 100
		go func(v *int) {
			*v = 100
		}(&x)

		// Set y to 5
		go func(v *int) {
			*v = 5
		}(&y)

		/*
		  A Race condition occurs when multiple threads are trying to access and manipulate the same variable.
		  The code below are all accessing and changing the same value.  Divide x (100) by y (5)
		  Except if y is not assigned 5 before x is assigned 100, y's initialized/zero value of 0 is used,
		  Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable,
		  which causes a divide by zero exception.
		*/
		go func(v1 int, v2 int) {
			fmt.Println(v1 / v2)
		}(x, y)

		i += 1

		fmt.Printf("%d\n", i)
	}
}
