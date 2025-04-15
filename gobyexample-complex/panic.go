package main

import "os"

func panicFn() {
	panic("a problem")

	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}
}
