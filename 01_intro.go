package main

import "fmt"

func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func testMax() {
	fmt.Println(Max(10, 100))
	fmt.Println(Max(9.9, 10.1))
}
