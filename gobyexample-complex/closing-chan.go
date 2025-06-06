package main

import "fmt"

func closingChan() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // j(val), more(bool, open or not)
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // to wait on this go routine
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	<-done
	// chan closed
	// jobs <- 6
	// fmt.Println(<-jobs)
}
