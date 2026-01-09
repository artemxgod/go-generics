package main

import "fmt"

func testLimits() {
	box := NewBox(10)
	fmt.Println(box)

	// cannot infer
	// emptyBox := NewBox()

	//explicit specification
	emptyBox := NewBox[int](0)
	fmt.Println(emptyBox)
}

type Box[T any] struct {
	value T
}

func NewBox[T any](value T) *Box[T] {
	return &Box[T]{value: value}
}
