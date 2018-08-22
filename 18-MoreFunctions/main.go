package main

import (
	"fmt"
)

// This is a function that returns another function
// Remember if returning a function if the returned function
// also has a return you need the include the return type in the type
// see the example below
func makeGreeter() func() string {
	return func() string {
		return "Hello World!"
	}
}

func main() {
	// Here we are setting the value of greater
	// to the value that is returned by the
	// makeGreeter function
	greeter := makeGreeter()

	// Here we are calling the function expression
	// "greeter" that we just created
	fmt.Println(greeter())
}
