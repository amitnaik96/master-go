package main

import "fmt"

// A closure is a function that "closes over" variables from its outer scope
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func counter() int {
	i := 0
	i++
	return i
}

func closure() {
	nxtInt := counter()
	fmt.Println(nxtInt)
	fmt.Println(nxtInt) // i same again 1

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt()) // i is remembered, Go keeps them alive
}
