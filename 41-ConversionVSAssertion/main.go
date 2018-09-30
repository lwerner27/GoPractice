package main

import (
	"fmt"
)

func main() {
	// Example of conversion
	var x = 12
	var y = 12.01245
	// We are using float64() to convert x (just an int) to a float64
	fmt.Println(y + float64(x))
	// We can also do the opposite and convert y to an int
	fmt.Println(x + int(y))

	// here is an example of assertion
	var name interface{} = "Lucas"

	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("name is not a string.")
	}
}
