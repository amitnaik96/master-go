package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working..")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSync() {
	done := make(chan bool, 1)
	go worker(done)

	// wait for channel to be notified
	<-done
}
