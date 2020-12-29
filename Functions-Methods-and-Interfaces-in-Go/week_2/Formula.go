// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("\nPlease enter acceleration: ")
// 	var acceleration float64
// 	fmt.Scanln(&acceleration)

// 	fmt.Printf("\nPlease enter velocity: ")
// 	var velocity float64
// 	fmt.Scanln(&velocity)

// 	fmt.Printf("\nPlease enter displacement: ")
// 	var displacement float64
// 	fmt.Scanln(&displacement)

// 	fmt.Printf("\nPlease enter time: ")
// 	var time float64
// 	fmt.Scanln(&time)

// 	fn := GenDisplaceFn(acceleration, velocity, displacement)
// 	fmt.Printf("The resulting displacement value s(t) is [%f]\n", fn(time))
// }

// func GenDisplaceFn(acc, iVel, iDis float64) func(float64) float64 {
// 	fn := func(t float64) float64 {
// 		return ((0.5 * acc * t * t) + (iVel * t) + iDis)
// 	}
// 	return fn
// }
package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Enter acceleration:")
	fmt.Scanln(&a)
	var v0 float64
	fmt.Print("Enter the initial velocity:")
	fmt.Scanln(&v0)
	var s0 float64
	fmt.Print("Enter initial displacement:")
	fmt.Scanln(&s0)
	var t float64
	fmt.Print("Enter time:")
	fmt.Scanln(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1/2*a*t*t + v0*t + s0
	}
}
