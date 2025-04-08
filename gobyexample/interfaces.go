package main

import "fmt"

// An interface in Go defines a set of method signatures, and any type that implements those methods satisfies the interface — no explicit declaration needed.

type geometry interface {
	area() int
	perim() int
}

type rect struct {
	l, b int
}

func (r rect) area() int {
	return r.l * r.b
}

func (r rect) perim() int {
	return 2 * (r.l + r.b)
}

// now this struct satifies the interface.

func measure(g geometry) {
	fmt.Println("area:", g.area())
	fmt.Println("premiter:", g.perim())
}

func interfaces() {
	r1 := rect{2, 3}
	measure(r1)
}
