package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func goroutine() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
}
