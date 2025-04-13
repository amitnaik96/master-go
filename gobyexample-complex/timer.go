package main

import (
	"fmt"
	"time"
)

func timerFn() {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C // blocks on timer channel
	fmt.Println("Timer fired!")

	// you can stop timer but not in time.SSleep
	timer2 := time.NewTimer(2 * time.Second)
	// done := make(chan bool)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
		// done <- true
	}()
	// <-done
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}
}
