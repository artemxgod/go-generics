package main

import "fmt"

// func duplication
func LastInt(s []int) int {
	return s[len(s)-1]
}

func LastString(s []string) string {
	return s[len(s)-1]
}

// interface way
func LastIf(s interface{}) interface{} {
	switch v := s.(type) {
	case []int:
		return v[len(v)-1]
	case []string:
		return v[len(v)-1]
	}

	return 0
}

// using generic
func Last[T any](s []T) T {
	return s[len(s)-1]
}
func testLast() {
	sInt := []int{1, 2, 3}
	sString := []string{"1", "2", "3"}

	// lose compile-time type checking
	val := LastIf(sString).(string)
	fmt.Println(val)

	fmt.Println(Last(sInt))
	fmt.Println(Last(sString))
}
