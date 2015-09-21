// -*- tab-width:4 -*-
// Exercise: Slices
//
//    https://tour.golang.org/moretypes/14

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dx)
	for x := range v {
		v[x] = make([]uint8, dy)
		for y := range v[x] {
//			v[x][y] = uint8((x + y)/2)
//			v[x][y] = uint8(x*y)
			v[x][y] = uint8(x^y)
		}
	}
	return v
}

func main() {
	pic.Show(Pic)
}
