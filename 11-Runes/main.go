package main

import (
	"fmt"
)

func main() {
	// ====================================================
	// Converts Hello Into Bytes
	// ====================================================
	fmt.Println([]byte("Hello"))

	// ====================================================
	// This will print out utf-8 values
	// ====================================================
	for i := 50; i <= 126; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	// ====================================================
	// This will be a rune because of the single quotes
	// ====================================================
	foo := 'a'

	// ====================================================
	// This will print out the decimal value of foo
	// ====================================================
	fmt.Println(foo)

	// ====================================================
	// This will show the type of variable foo
	// ====================================================
	fmt.Printf("%T\n", foo)
}
