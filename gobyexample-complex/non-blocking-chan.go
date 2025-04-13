package main

import "fmt"

// sends , receives from chan -> blocking
// select, default -> non blocking

func nonBlockingChan() {
	messages := make(chan string)

	messages <- "msg"
	fmt.Println(<-messages)

	select {
	case messages <- "msg": // blocked forever. refer img
		fmt.Println("sent msg")
	default:
		fmt.Println("no message sent")
	}
}
