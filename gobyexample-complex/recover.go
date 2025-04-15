package main

import "fmt"

func myPanic() {
	panic("yo")
}

func recoverFn() {

	// called when enclosing func(revoverFn) panics. reover() -> built in
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered, err:", r)
		}
	}()

	myPanic()
}
