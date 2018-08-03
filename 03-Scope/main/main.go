package main

import (
	"fmt"

	"github.com/lwerner27/03-Scope/test"
)

// This variable (age) has package level scope
var age = 25

func main() {
	// This variable (name) has block level scope
	name := "Lucas Werner"
	fmt.Println(name)
	fmt.Printf("Type of age variable: %T \n", age)
	foo()
	test.PrintFirstName()

	// I cannot call print last name in this file
	// It has not been exported visibly

}

// This function has package level scope
func foo() {
	fmt.Println("Foo")
}
