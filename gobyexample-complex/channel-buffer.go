package main

import "fmt"

func bfferChannel() {
	messages := make(chan string, 2) // capacity 2

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
