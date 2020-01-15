package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, init_velocity, init_displacement float64) func(float64) float64 {
	return func(time float64) float64 {
		displacement := (1 / 2 * (acceleration * math.Pow(time, 2))) + (init_velocity * time) + init_displacement
		return displacement
	}
}

func main() {
	var fn_acc, fn_vel, fn_disp float64
	fmt.Println("Enter Acceleration:")
	fmt.Scan(&fn_acc)
	fmt.Println("Enter Initial Velocity:")
	fmt.Scan(&fn_vel)
	fmt.Println("Enter Initial Displacement:")
	fmt.Scan(&fn_disp)
	fn := GenDisplaceFn(fn_acc, fn_vel, fn_disp)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
