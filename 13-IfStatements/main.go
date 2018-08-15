package main

import (
	"fmt"
)

func main() {
	test := false
	num := 234

	// Simplest if statement
	// If true do one thing else do something different
	if test {
		fmt.Println("The value of test is true.")
	} else {
		fmt.Println("The value of test is false.")
	}

	// In this example the ! will reverse the value of the variable
	// Meaning the main if response will only run if test is false
	if !test {
		fmt.Println("This will only run if test is false.")
	} else {
		fmt.Println("This will only run if test is true.")
	}

	// This is an example of using an initalization statement in a if statement
	// This can help us keep a tight scope.
	// In this example test2 will not be accessible outside of the if statement
	if test2 := test; test2 {
		fmt.Println("Test2 is true which means that test is true.")
	} else {
		fmt.Println("Test2 is false which means that test is false.")
	}

	// This is an example of an if else statement with an else if
	if num == 1 {
		fmt.Println("The number in the num variable equals 1.")
	} else if num == 2 {
		fmt.Println("The number in the num variable equals 2.")
	} else if num == 3 {
		fmt.Println("The number in the num variable equals 3.")
	} else {
		fmt.Println("Not sure what number is in the number variable.")
	}

}
