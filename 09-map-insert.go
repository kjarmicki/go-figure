package main

import "fmt"

func MapInsert() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
		3: false,
		4: true,
	}

	for k, v := range m {
		if v == true {
			m[k+10] = true
		}
	}

	fmt.Println(len(m))
}

/* 
 * Values inserted into map during iteration may or may not be iterated over. There's no way to enforce either
 * iteration or skipping for newly inserted values.
 */