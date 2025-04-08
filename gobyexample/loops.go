package main

import "fmt"

func loops() {
	i := 0
	for i <= 3 {
		fmt.Printf("%v ", i)
		i++
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("%v ", i)
	}

	for {
		fmt.Println("yo")
		break
	}

	for i := range 6 {
		fmt.Printf("%v ", i)
	}
}
