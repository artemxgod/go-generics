package main

import "fmt"

func testStructures() {
	intStack := Stack[int]{}
	intStack.Push(42)
	intStack.Push(100)
	val, ok := intStack.Pop() // val is of type int, no assertion needed
	if ok {
		fmt.Println(val) // 100
	}

	strStack := Stack[string]{}
	strStack.Push("Hello")
	strStack.Push("Kain")
	top, _ := strStack.Pop() // top is of type string
	fmt.Println(top)         // "world"

	p1 := Pair[int, string]{First: 42, Second: "answer"}
	fmt.Printf("%d is the %s\n", p1.First, p1.Second)

	p2 := Pair[string, bool]{First: "enabled", Second: true}
	fmt.Println(p2.First, "=", p2.Second)
}

// stack is LIFO containter
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // zero val of type T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Pair holds two values of possibly different types
type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}
