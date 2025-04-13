package main

import (
	"fmt"
	"time"
)

func ratelimitingFn() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	// 1 req every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req)
	}
}
