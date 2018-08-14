package main

import (
	"fmt"
)

// SwitchOnType requires the use of an interface
// We will talk more about this later
func SwitchOnType(x interface{}) {
	switch x.(type) { // This is an assertion; asserting, "x is of this type"
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("Can not identify type of x")
	}
}

func main() {
	// Unlike other languages Go does not require a break
	switch "Dave" {
	case "Lucas":
		fmt.Println("Hey Lucas.")
	case "Dave":
		fmt.Println("Hey Dave.")
		// This makes the next entry also evaluate as true
		fallthrough
	case "Tom":
		fmt.Println("Hey Tom.")
	default:
		fmt.Println("Sorry, my mom told me not to talk to strangers.")
	}

	SwitchOnType(12)
	SwitchOnType("Lucas Werner")

}
