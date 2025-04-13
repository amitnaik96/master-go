package main

import "fmt"

func firstElement[T any](s []T) {
	fmt.Println(s[0])
}

func generics() {
	nums := []int{35, 67, 12}
	strs := []string{"yo", "hello"}

	// firstElement[int](nums)
	firstElement(nums)
	firstElement(strs)
}
