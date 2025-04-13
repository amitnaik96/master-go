package main

import (
	"fmt"
	"time"
)

func selectFunc() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2"
	}()

	// select exits when one case is executed.
	// select {
	// case msg1 := <-c1:
	// 	fmt.Println(msg1)
	// case msg2 := <-c2:
	// 	fmt.Println(msg2)
	// }

	// wrap in loop if you want both.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
