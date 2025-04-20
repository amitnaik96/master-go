package main

import (
	"fmt"
	"package-go/helper"
)

func main() {
	greet("morning")

	helper.Greetme()
	fmt.Println(helper.Name)
}

// go run main.go  -> won't work

// works
// go run main.go helper.go
// go run .
