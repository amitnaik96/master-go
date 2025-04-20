package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go")
}

func httpServerFn() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":5000", nil)
}
