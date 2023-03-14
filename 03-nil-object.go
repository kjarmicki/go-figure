package main

import "fmt"

type Greeter struct{}

func (g *Greeter) sayHi() {
	fmt.Println("hi there")
}

func NilStruct() {
	var g *Greeter
	if g == nil {
		g.sayHi()
	}
}

/*
 * Struct methods are kind of a syntax sugar. They do not strictly belong to a particular struct.
 * This code could've been written as:
 *
 *	type Greeter struct{}
 *
 *	func sayHi(g *Greeter) {
 *		fmt.Println("hi there")
 *	}

 *	func NilStruct() {
 *		var g *Greeter
 *		if g == nil {
 *			sayHi(g)
 *		}
 *	}
 */
