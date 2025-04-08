package main

import "fmt"

func fact(n int) int {
	if n <= 1 {
		return 1
	}

	return n * fact(n-1)
}

func recursive() {
	fmt.Println(fact(5))

	// using anonymous functions

	fact := func(n int) int {
		if n <= 1 {
			return n
		}

		return n * fact(n-1)
	}

	fmt.Println(fact(5))
}
