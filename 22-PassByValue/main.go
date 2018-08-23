package main

import "fmt"

func main() {
	myAge := 26
	fmt.Println("My Age is: ", myAge)

	// This is how we pass addresses
	fmt.Println("The address where my age is stored is: ", &myAge)

	// Here we will change the value of myAge using a function
	changeAge(&myAge)

	fmt.Println("My new age is: ", myAge)

}

// This function takes one argument
// That argument should be the memory address of an int
// This will give us the ability to change the value
// of a variable from a different scope
func changeAge(z *int) {
	// This will print the value stored at z
	fmt.Println("This is the value of z: ", *z)

	// This will print the memory address of z
	// This should be the same as the memory address for myAge
	fmt.Println("This is the address where data for z is stored: ", z)

	// This is how we can change the value of z and myAge at the same time
	*z = 27

	fmt.Println("Your age has been updated.")
}

// Everything in Go is pass by value.
