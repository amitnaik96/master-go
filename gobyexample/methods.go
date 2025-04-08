package main

import "fmt"

// class like structure as in C++, ..

type Rect struct {
	l, b int
}

// method
func (r *Rect) area() int { // reciever r
	return r.l * r.b
}

func methods() {
	r := Rect{10, 5}

	fmt.Println(r)
	fmt.Println(r.area())
}
