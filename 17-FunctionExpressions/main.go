package main

import (
	"fmt"
)

// This is and example of a function declaration
func declaration() {
	fmt.Println("This is a function declaration.")
}

func main() {

	// This is an example of a function expression
	// This is the only way to create a function inside another function
	expression := func() {
		fmt.Println("This is a function expression.")
	}

	declaration()
	expression()

}
