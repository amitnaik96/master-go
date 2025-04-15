package main

import (
	"encoding/json"
	"fmt"
)

func jsonFn() {
	name := "Amit"
	nameJson, _ := json.Marshal(name)
	fmt.Println(string(nameJson)) // string() else in byte[]

	var name2 string
	json.Unmarshal(nameJson, &name2)
	fmt.Println(name2)

	fruits := []string{"peach", "apple"}
	fruitsJson, _ := json.Marshal(fruits)
	fmt.Println(string(fruitsJson))

	scores := map[string]int{"Alice": 30, "bob": 28}
	scoresJson, _ := json.Marshal(scores)
	fmt.Printf(string(scoresJson))
}
