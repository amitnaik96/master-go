package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	var name string

	fmt.Println("Enter name")
	fmt.Scan(&name)

	wg.Add(1)        // 1 thread added to wait
	go display(name) // new thread spawned here.

	wg.Wait()
}

func display(name string) {
	time.Sleep(10 * time.Second)
	fmt.Printf("your name is %v\n", name)

	wg.Done() // thread execution done remove from wg.Add
}
