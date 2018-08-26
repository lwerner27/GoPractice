package main

import "fmt"

func main() {
	// Example of an array, notice the 2 in side the brackets
	// This sets arrays apart from slices, the 2 indicates this array will have 2 items
	// arrays aren't dynamic length but slices are.
	var x [10]int

	fmt.Println()
	fmt.Println("This is x before we have updated the values of the array:")
	fmt.Println(x) // This should print: [0,0,0,0,0,0,0,0,0,0]

	fmt.Println()

	// Here we are updating the values of the array
	// by targeting a specific index and setting it equal to our value.
	for i := 0; i < 10; i++ {
		x[i] = i
	}

	fmt.Println("This is x after we have updated the values of the array:")
	fmt.Println(x) // This should print: [0,1,2,3,4,5,6,7,8,9]
	fmt.Println()
}
