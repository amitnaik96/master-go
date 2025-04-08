package main

import "fmt"

func switchFunc() {
	age := 19

	// alternative to if-else
	switch {
	case age > 18:
		fmt.Println("drive")
	case age <= 18:
		fmt.Println("no way")
	}
}
