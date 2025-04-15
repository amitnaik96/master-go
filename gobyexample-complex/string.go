package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func stringFn() {
	p(s.Contains("text", "ex"))
	p(s.Count("text", "t"))
	p(s.Join([]string{"a", "b"}, "c"))
	p(s.Repeat("a", 5))
	p(s.ToLower("TEST"))
	// .. many more fns
	// refer string formating..
}
