package main

import (
	"fmt"
	"iter"
)

// iter.Seq[int] -> iterator of sequence values
// yield -> anonymous func
// the below two lines are function signatures to return an iterator. you can range(for) on them
func count() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func rangeOverIterators() {
	for num := range count() {
		fmt.Println(num)
	}
}
