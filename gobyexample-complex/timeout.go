package main

import (
	"fmt"
	"time"
)

func timeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "c1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // this gets returned after 1 sec
		fmt.Println("timeout!")
	}
}
