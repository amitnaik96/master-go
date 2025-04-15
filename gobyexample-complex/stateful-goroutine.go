package main

import "fmt"

func statefulGoroutine() {
	type readReq struct {
		key  int
		resp chan int
	}

	state := map[int]int{1: 35}

	reads := make(chan readReq)

	// single goroutine that own the state
	go func() {
		for {
			req := <-reads
			req.resp <- state[req.key]
		}
	}()

	respChan := make(chan int)
	reads <- readReq{key: 1, resp: respChan}
	val := <-respChan

	fmt.Println(val)
}
