package test

import (
	"fmt"
)

// PrintFirstName function as package level scope
// This function is being visibly exported
func PrintFirstName() {
	fmt.Println("Lucas")
}

// This function also has package level scope
// But it is not being exported visibly
// Unlike PrintFirstName it can not be used in main.go
func printLastName() {
	fmt.Println("Werner")
}
