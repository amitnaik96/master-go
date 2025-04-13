package main

import "fmt"

type argError struct {
	arg     int
	message string
}

// implement error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 45 {
		return -1, &argError{arg, "No 45"}
	}
	return 2, nil
}

func customError() {
	_, e := f(45)
	fmt.Println(e)
}
