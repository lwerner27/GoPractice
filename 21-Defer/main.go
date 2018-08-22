package main

import (
	"fmt"
)

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World!")
}

func main() {
	// the defer keyword makes whatever it is attached to execute
	// right before the containing function finishes its execution.
	// So in this example right before the main func finishes running
	// it will call the "world" fucntion
	defer world()
	hello()
}

// You can have multiple defer statements inside of one function
