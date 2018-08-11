package main

import (
	"fmt"
)

func main() {
	// Most Basic for loop example
	// Prints nummbers 0 through 9
	// Basically the exact same syntax as javascript
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Nested Loops
	for i := 0; i < 10; i++ {
		fmt.Println("This is the outer loop: ", i)
		for j := 0; j < 3; j++ {
			fmt.Println("This is the INNER loop: ", j)
		}
	}
}
