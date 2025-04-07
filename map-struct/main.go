package main

import (
	"fmt"
	"strconv"
)

func main() {

	// map
	var userData = make(map[string]string) // key-value pair map

	name := "amit"
	tkt := 15

	userData["name"] = name

	// No mixed data type in Go. convert int to string
	userData["tickets"] = strconv.FormatInt(int64(tkt), 10)

	fmt.Println(userData)

	// list(slice) of maps
	var bookings = make([]map[string]string, 0) // give initial size

	bookings = append(bookings, userData)

	for _, booking := range bookings {
		fmt.Printf("Name %v, noOfTickets: %v", booking["name"], booking["tickets"])
	}

	// struct -> mixed data type
	type UserData struct {
		name string
		age  int
	}

	userData2 := UserData{
		name: "amit",
		age:  22,
	}
	fmt.Println(userData2.name)

	bookings2 := []UserData{}
	bookings2 = append(bookings2, userData2)
	fmt.Println(bookings2)
}
