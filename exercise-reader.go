// -*- tab-width:4 -*-
// Exercise: Readers
//
//    https://tour.golang.org/methods/11

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (p MyReader) Read(buf []byte) (int, error) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 'A'
	}
	return len(buf), nil
}

func main() {
	reader.Validate(MyReader{})
}
