package main

import (
	"cmp"
	"fmt"
)

func testConstraint() {
	fmt.Println(MaxC(int8(10), -10))
	fmt.Println(Square(12342))
}

func MaxC[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// custom constraints

// requires a String() method
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

// A constraint for numeric types that support multiplication
type Numeric interface {
	~int | ~int64 | ~float64 // Using '~' means "underlying type is..."
}

func Square[T Numeric](x T) T {
	return x * x
}
