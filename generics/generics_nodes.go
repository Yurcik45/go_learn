package main

// List represents a singly-linked list that holds
// values of any type.

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	test := List[int]{nil, 2}
	fmt.Println("test: ", test)
	test1 := List[int]{&test, 3}
	fmt.Println("test1: ", test1)
}
