package main

import (
	"fmt"
	"strconv"
)

func parseNumber() {
	f, _ := strconv.ParseFloat("12.34", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("35", 0, 64)
	fmt.Println(i)

	k, _ := strconv.Atoi("56")
	fmt.Println(k)
}
