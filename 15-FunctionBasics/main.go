package main

import (
	"fmt"
)

// Most basic example of function declaration
// showName is the name of the function
// This fucntion requires 1 parameter
// The parameter is the name in side of the parens
// It needs to be a string
func showFirstName(fName string) {
	fmt.Println(fName)
}

// This fuction has a return
// The return type is stated between the parens and the curly braces
func showFullName(fName, lName string) string {
	return fmt.Sprint(fName, lName)
}

func main() {
	showFirstName("Lucas")                       // Here is an example of an arguement
	showFullName("Lucas", "Werner")              // Notice thing is printed here
	fmt.Println(showFullName("Lucas", "Werner")) // We needed to print the return value
}

// Parameters are in the function declaration
// Arguements are used where the function is called

// Unlike most other languages you can do multiple returns
