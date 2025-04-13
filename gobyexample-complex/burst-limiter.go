package main

import (
	"fmt"
	"time"
)

func burstLimiterFn() {
	burstLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t // blocked until there is space
		}
	}()

	burstRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequest <- i
	}
	close(burstRequest)
	for req := range burstRequest {
		<-burstLimiter
		fmt.Println("request", req)
	}
}

// first burstLimiter had 3tokens.
// then every 200ms 1 token(time)
// that is why the first three in burstRequest burst out. and the later ones come every 200ms
// first 3 req gets burst out.
