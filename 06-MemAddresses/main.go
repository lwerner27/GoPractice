package main

import (
	"fmt"
)

func main() {
	a := 43

	// Shows the assigned value of variable a
	fmt.Println("a - ", a)

	// Shows the memory addres of variable a
	fmt.Println("a's memory addresss - ", &a)

	// Show the memory address for variable a in decimal
	fmt.Printf("%d\n", &a)
}
