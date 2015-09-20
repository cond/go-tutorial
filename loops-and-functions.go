// -*- tab-width:4 -*-
// Exercise: Loops and Functions
//
//    https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x)
	z0 := float64(x+10)
	n := 0
	for math.Abs(z - z0) > 0.000000001 {
		z0 = z
		z = z - (z * z - x)/(2 * z)
		n++
	}
	fmt.Println(n, " times")
	return z
}

func main() {
	for i := 2; i <= 20; i++ {
		fmt.Println("sqrt(", i, ") = ", Sqrt(float64(i)))
	}
}
