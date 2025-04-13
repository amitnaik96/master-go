package main

import "fmt"

func channels() {
	messages := make(chan string) // create channel message

	go func() {
		messages <- "ping" // goroutine sends value
	}()

	// fmt.Println(<-messages)
	msg := <-messages // main goroutine receives it
	fmt.Println(msg)
}
