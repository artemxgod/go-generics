package main

import "fmt"

func testFlexibility() {
	// Convert []int to []string
	numbers := []int{1, 2, 3, 4}
	strings := Map(numbers, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println(strings) // [#1 #2 #3 #4]

	// Double each float64 value
	floats := []float64{11.1, 22.2, 33.3}
	doubled := Map(floats, func(f float64) float64 {
		return f * 2
	})
	fmt.Println(doubled) // [22.2 44.4 66.6]

	// get even numbers
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evens) // [2, 4, 6]

	// get words longer than 3 letters
	words := []string{"go", "generic", "functions", "are", "cool"}
	longWords := Filter(words, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println(longWords) // [generic functions cool]
}

// gets a slice of type T changes it according to the transform function and returns a slice of type U
func Map[T, U any](input []T, transform func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = transform(v)
	}
	return result
}

// Filter returns elements that satisfy a predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
