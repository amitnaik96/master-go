package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup

func worker2(id int) {
	// defer wg.Done()
	fmt.Printf("worker %d start\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d end\n", id)
}

func waitFn() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()
}
