package main

import (
	"fmt"
	"os"
)

func readFile() {
	// dir, _ := os.Getwd()
	// fmt.Println(dir)

	dat, err := os.ReadFile("./tmp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dat))
}
