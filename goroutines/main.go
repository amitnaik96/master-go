package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	var email string
	var noTkt int

	for {
		fmt.Println("Enter name, email, and no of tickets")
		fmt.Scan(&name, &email, &noTkt)

		// sendTicket(name, email, noTkt) // blocks main thread for 10 secs
		go sendTicket(name, email, noTkt) // go : Now this runs on a seperate thread.
		// Note: here we have infinte loop so this new thread returns. What if the main thread exits. refer wait.go
	}
}

func sendTicket(name string, email string, tkts int) {
	time.Sleep(10 * time.Second) // simulating delay
	var ticket = fmt.Sprintf("Tickets %v, Email %v", tkts, email)
	fmt.Printf("Here is your ticket %v : %v", name, ticket)
}
