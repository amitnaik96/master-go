package main

import "fmt"

// any no. of argumnents
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func varadicFun() {
	r1 := sum(1, 2)
	r2 := sum(1, 2, 3)

	fmt.Println(r1, r2)
}
