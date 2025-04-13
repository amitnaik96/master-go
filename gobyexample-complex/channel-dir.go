package main

import "fmt"

// chan<- (recieving chan) , string
func ping(pings chan<- string, msg string) {
	pings <- msg // receives
}

// <-chan sending channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDir() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed msg")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
