package main

import "fmt"

type FriendlyPerson struct {
	Name        string
	FriendNames []string
}

func FriendlyPersonEquals() {
	c1 := FriendlyPerson{Name: "Alice", FriendNames: []string{"Bob"}}
	c2 := FriendlyPerson{Name: "Alice", FriendNames: []string{"Bob"}}
	// fmt.Println(c1 == c2)
	fmt.Println(c1, c2)
}

/*
 * Two structs can be compared if all of their values are comparable. Slices, maps and functions are not comparable.
 */
