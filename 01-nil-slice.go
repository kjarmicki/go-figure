package main

import "fmt"

func NilSlice() {
	var s []string
	if s == nil {
		s = append(s, "a")
	}
	fmt.Println(s)
}

/*
 * This is a nil slice. It's similar to an empty slice but it's also equal to nil, despite having a type.
 * How to create an empty slice? make([]string, 0)
 * Contrary to an empty slice, it doesn't require any allocation because it has no backing array.
 * Use case: a function returning a slice of elements of unknown length which can be zero.
 */
