package main

import (
	"fmt"
)

func main() {
	// Example of or expression (same syntax as JavaScript)
	if true || false {
		fmt.Println("This will run as long as one of the two values is true.")
	}

	if true && false {
		fmt.Println("This will only run if both values are true")
	} else {
		fmt.Println("This will run if either of the values is false.")
	}
}
