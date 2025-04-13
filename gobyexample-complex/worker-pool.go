package main

import (
	"fmt"
	"time"
)

func worker1(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("worker %d is processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func workerPoolFn() {
	jobs := make(chan int)

	for w := 1; w <= 2; w++ {
		go worker1(w, jobs)
	}

	for j := 1; j <= 3; j++ {
		jobs <- j
	}

	time.Sleep(3 * time.Second)
}
