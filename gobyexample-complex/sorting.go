package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sortingFn() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println(strs)

	sortingByLen()
}

func sortingByLen() {
	fruits := []string{"peach", "kiwi"}

	lenComp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenComp)
	fmt.Println(fruits)
}
