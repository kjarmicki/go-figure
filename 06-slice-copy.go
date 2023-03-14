package main

import "fmt"

func SliceCopy() {
	dst := make([]int, 3)
	src := []int{1, 2, 3}
	copy(dst, src)
	fmt.Println(dst)
}

/* The amount of copied elements is either source's or destination's length, whichever is smaller. */
