package main

import (
	"fmt"
)

// This is a typed constant
const name string = "Lucas Werner"

// This is an untyped constant, also known as a kind
// These are not yet forced to follow the rules that prevent
// The combining of of differently typed values
const hello = "Hello World"

func main() {
	fmt.Println("Hello World")
}
