package main

import (
	"fmt"
)

type CountError struct {
	Count int
}

func (m *CountError) Error() string {
	return fmt.Sprintf("there were %d errors along the way", m.Count)
}

func validateString(foo string) error {
	var m *CountError
	if foo != "foo" {
		m.Count += 1
	}
	return m
}

func CountErrorValidate() {
	err := validateString("foo")
	if err != nil {
		fmt.Printf("string is invalid: %v", err)
	}
}

/*
 * Interfaces are dispatch wrappers (boxes). They consist of two things: a type T and a value V.
 * So here we have (T=*CountError V=nil).
 * An interface is equal to nil ONLY IF its type is equal to nil, value is irrelevant.
 * How do you make an interface's type equal to nil? Return nil explicitly.
 *	func validateString(foo string) error {
 *		m := &CountError{}
 *		if foo != "foo" {
 *			m.Count += 1
 *		}
 *		if m.Count > 0 {
 *			return m
 *		}
 *		return nil -> (T=nil V=nil)
 *	}
 */
