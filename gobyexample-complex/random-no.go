package main

import (
	"fmt"
	"math/rand/v2"
)

func randomInt() {
	fmt.Println(rand.IntN(100)) // 0 - 100
	fmt.Println(rand.Float64()) // 0 - 1.0
}
