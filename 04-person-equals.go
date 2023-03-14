package main

import "fmt"

type Person struct {
	Name string
}

func PersonEquals() {
	c1 := Person{Name: "Alice"}
	c2 := Person{Name: "Alice"}
	fmt.Println(c1 == c2)
}

/* As opposed to objects in a lot of other languages, Go structs are compared by value. */
