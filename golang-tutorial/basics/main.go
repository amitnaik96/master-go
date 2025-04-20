package main

import (
	"fmt"
	"strings" // format package
)

func main() {
	// fmt.Print("Hello World")
	fmt.Println("Hello World") // new line

	var name = "Amit" // error if u don't use this var
	fmt.Println("My name is", name)

	const tickets = 50
	var remainingTickets = tickets

	fmt.Println("Remaining tickets", remainingTickets)

	fmt.Printf("%v tickets are remaining\n", remainingTickets)

	// data types
	var userName string
	var userTickets int
	userName = "Amit"
	fmt.Println(userName)
	fmt.Println(userTickets)

	// get types
	fmt.Printf("tickets %T\n userName %T", tickets, userName)

	// var tkt uint
	// tkt = -1 cannot bcoz uint -> unsigned int

	// better syntax
	age := 2 // only works for var (cannot explicitly delcare data-type)
	fmt.Print(age)

	// pointers
	fmt.Println(&remainingTickets)

	// user input scan
	var user string
	// fmt.Scan(&user) // give memory addr -> pointer
	fmt.Printf("Your name is %v\n", user)

	// arrays
	// var bookings = [10]string{"Amit", "Robert"}
	var bookings [10]string // delcare array type

	bookings[0] = "amit"
	bookings[1] = "john"

	fmt.Println(bookings)
	fmt.Println(bookings[0])
	fmt.Println(len(bookings))

	// slices -> dynamic arrays
	// var bkngs []string
	bkngs := []string{}
	bkngs = append(bkngs, "amit")
	fmt.Println(bkngs)

	// loops -> for loop with types

	//infinite loop
	// for {
	// fmt.Println("Hello mate!") // ctrl + c
	// }

	// for each loop
	names := []string{"amit naik", "john doe"}
	firstNames := []string{}

	// for index, name := range names {  // actually return index, element
	for _, name := range names { // _ for unused vars
		var fName = strings.Fields(name) // splits strings with " "
		firstNames = append(firstNames, fName[0])
	}
	fmt.Println(firstNames)

	// if statement
	if remainingTickets == 0 {
		fmt.Println("Bookings closed")
	}

	// if else statement
	if age >= 18 {
		fmt.Println("18+")
	} else {
		fmt.Println("Grow up kid!")
	}

	// else if statement
	if age >= 18 {
		fmt.Println("18+")
	} else if age == 17 {
		fmt.Println("wait an year mate")
	} else {
		fmt.Println("Grow up kid!")
	}

	// conditional for
	for age < 18 {
		age++
		fmt.Println(age)
	}

	// user input validation
	email := "amit@gmail.com"
	isValidEmail := strings.Contains(email, "@") && len(email) >= 2
	if isValidEmail {
		fmt.Println("Valid email!")
	}
	// && , ||, ! same as other langs

	// switch statement
	city := "London"
	switch city {
	case "New York": // code here
	case "London": // code here
	case "Berlin", "Mexico": // code here
	default:
		fmt.Println("Default")
	}
}
