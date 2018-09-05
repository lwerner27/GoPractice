package main

import (
	"fmt"
)

func main() {
	myGreetings := map[string]string{
		"morning":   "Good morning.",
		"afternoon": "Good afternoon.",
		"evening":   "Good evening.",
	}

	// In this value "key" and "value" are just variable names
	// for all intensive purposes they could be "k" and "v"
	for key, value := range myGreetings {
		fmt.Println(key, " - ", value)
	}
}
