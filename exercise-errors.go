// -*- tab-width:4 -*-
// Exercise: Errors
//
//    https://tour.golang.org/methods/9

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: Negative number %v is given.", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(x)
	z0 := float64(x + 10)
	for math.Abs(z-z0) > 0.000000001 {
		z0 = z
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
