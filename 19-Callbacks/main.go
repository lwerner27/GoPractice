package main

import (
	"fmt"
)

// This is our function that calls a callback function
// this is very similar to how callbacks work in JavaScript
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	// This is our callback function in action
	// again this is very similar to the JavaScript notation
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}
