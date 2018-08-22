package main

import "fmt"

// This is an exampel of a recursive function
// Recursive functions call themselves
func factorial(x int) int {
	if x == 0 {
		return 1
	}

	// This is where our recursion is happening
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}

// Most of the time recusion has a significant performance cost
// Which rarely makes it the best choice
