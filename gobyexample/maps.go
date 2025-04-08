package main

import "fmt"

func mapFunc() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 45

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("%v %v\n", k, v)
	}

	delete(m, "k2")
	clear(m)

}
