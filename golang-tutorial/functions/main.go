package main

import (
	"fmt"
	"strings"
)

// package level variables
var number = 45 // No :=

func main() {
	greet()
	greet2("evening")

	text := greet3()
	fmt.Println(text)

	name := "amit"
	email := "amit@gmail.com"

	isValidName, isValidEmail := validate(name, email)
	if isValidName && isValidEmail {
		fmt.Println("valid user!")
	}

	fmt.Println(number)
	check()
}

func greet() {
	fmt.Println("Good morning!")
}

func greet2(period string) {
	fmt.Printf("good %v!\n", period)
}

func greet3() string { // return type
	return "good night!"
}

// mutiple return types
func validate(name string, email string) (bool, bool) {
	isValidName := len(name) > 2
	isValidEmail := strings.Contains(email, "@")

	return isValidName, isValidEmail
}

func check() {
	fmt.Println(number)
}
