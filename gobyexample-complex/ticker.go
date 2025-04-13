package main

import (
	"fmt"
	"time"
)

func tickerFn() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	done <- true
}
