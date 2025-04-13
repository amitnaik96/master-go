package main

import (
	"errors"
	"fmt"
)

// last return val -> error convention
func ef1(arg int) (int, error) {
	if arg == 45 {
		return -1, errors.New("no 45")
	}
	return 2, nil
}

var error45 = fmt.Errorf("error variable")

func ef2(arg int) error {
	if arg == 45 {
		// return error45
		return fmt.Errorf("error : %w", error45)
	}
	return nil
}

func errorsFunc() {
	//errors are ignored by default. you need to display them like this!
	_, e := ef1(45)
	fmt.Println(e)

	_, e = ef1(4)
	fmt.Println(e)

	er := ef2(45)
	fmt.Println(er)
}
