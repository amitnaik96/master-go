package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock() // will run at end of fn
	c.counters[name]++
}

func mutexFn() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)

	wg.Wait()

	fmt.Println(c.counters)
}
