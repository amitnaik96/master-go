package main

import (
	"fmt"
	"slices"
)

func sliceFunc() {
	var nums [10]int // array
	nums[0] = 1
	fmt.Println(nums) // rest 9 -> 0, if string then ' '
	fmt.Println(nums[0])

	var numbers []int // slice
	fmt.Println(numbers)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)

	fmt.Println(numbers)
	fmt.Println(numbers[0])

	slices2()
}

func slices2() {
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	if slices.Equal(s, c) {
		fmt.Println("equal")
	}

	sl1 := s[0:1] // 0,1 (exclude 2)
	sl2 := s[1:]  // from 1 get all
	sl3 := s[:1]  // before 1 get all
	fmt.Println(sl1, sl2, sl3)

	twoD := make([][]int, 2)
	for i := 0; i < 2; i++ {
		twoD[i] = make([]int, 2)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = 7
		}
	}

	fmt.Println(twoD)
}
