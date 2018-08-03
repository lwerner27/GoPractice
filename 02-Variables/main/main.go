package main

import (
	"fmt"
)

// You can declare variables with their zero values
// To do this we use var and state what type of date the variable will hold
var a string
var b int
var c bool

func main() {
	// This is an example of short hand variable daclaration
	// This kind of declaration can only be done inside of a function
	name := "Lucas Werner"
	age := 25
	canDrink := true

	fmt.Println()
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Age: %v \n", age)
	fmt.Printf("Can Drink: %v \n", canDrink)
	fmt.Println()
	fmt.Printf("Name is a: %T \n", name)
	fmt.Printf("Ageis a: %T \n", age)
	fmt.Printf("CanDrink is a: %T \n", canDrink)
	fmt.Println()
}
