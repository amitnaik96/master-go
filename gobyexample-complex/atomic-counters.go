package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicCounterFn() {
	var ops atomic.Uint64

	// ops.Store(12)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops", ops.Load())
}
