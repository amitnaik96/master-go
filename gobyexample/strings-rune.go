package main

import (
	"fmt"
	"unicode/utf8"
)

// rune -> character

func stringRune() {
	s := "hello"

	fmt.Printf("%T\n", s)
	fmt.Printf("len : %v\n", len(s))                    // this does not work outside ascii chars
	fmt.Printf("len : %v\n", utf8.RuneCountInString(s)) // recommended

	var c rune = 'a'
	fmt.Printf("%T\n", c)
	fmt.Println(c)
	fmt.Printf("%c\n", c)
}
