package main

import (
	"fmt"
	"os"
)

func deferFn() {
	f := createFile("./tmp.txt")
	defer closeFile(f) // executes at last
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
