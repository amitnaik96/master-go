package main

import "fmt"

type base struct {
	num int
}

func (b base) display() {
	fmt.Println(b.num)
}

//embed base in container
type container struct {
	base
	str string
}

func structEmbed() {
	c := container{
		base: base{
			num: 1,
		},
		str: "yo",
	}

	fmt.Println(c)
	c.display()
}
