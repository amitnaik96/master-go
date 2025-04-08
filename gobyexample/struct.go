package main

import "fmt"

type Person struct {
	name string
	age  int
}

func structFunc() {
	fmt.Println(Person{"amit", 22})

	user := struct { // anonymous struct type
		username string
		admin    bool
	}{"amit", true}

	user.username = "random"
	fmt.Println(user)
}
