package main

import "fmt"

func main() {
	m := make([]string, 1, 25)            // This will make and empty slice
	fmt.Println("Curent value of m: ", m) // This will print and empty slice
	changeValue("Lucas", m)               // This will update the first value in the slice
	fmt.Println("Curent value of m: ", m) // This will print the updated value of the slice

}

// We do not need a memory addres for a slice because a slice is reference type
// This means just passing in the slice will be enough to update the slice throughout the code

func changeValue(y string, z []string) {
	// This will update the value at the zero index
	z[0] = y
	fmt.Println("The string has been added.")
}
