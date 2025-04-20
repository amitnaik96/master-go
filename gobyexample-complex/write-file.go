package main

import (
	"fmt"
	"os"
)

func writeFileFn() {
	d1 := []byte("hello\nthere")
	err := os.WriteFile("./tmp2.txt", d1, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
