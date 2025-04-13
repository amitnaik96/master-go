package main

import "fmt"

type Status int

const (
	Success     Status = 200
	NotFound    Status = 404
	ServerError Status = 500
)

// String built-in from Stringer interface
func (s Status) String() string {
	switch s {
	case Success:
		return "success"
	case NotFound:
		return "not-found"
	case ServerError:
		return "server-error"
	default:
		return "unknown"
	}
}
func enums() {
	fmt.Println(Success)
	fmt.Println(NotFound)
}
