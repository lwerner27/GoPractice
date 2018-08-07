package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println("This is the value of a: ", a)
	fmt.Println("This is the memory address of a: ", &a)

	// This line makes b a pointer to a's memory address
	// Points to an address where an int is stored
	var b = &a

	fmt.Println("This is b: ", b)
	// Example of dereferencing a variable
	fmt.Println("The value of a through b: ", *b)

	// This line says change the stored value at the address stored in b
	*b = 42
	fmt.Println("This is the new value of a: ", a)

	// Everything in Go is PASS BY VALUE

}
