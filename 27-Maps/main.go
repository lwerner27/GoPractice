package main

import (
	"fmt"
)

func main() {
	// Creates the Map very simlar to slices
	// the "string" in the example says our keys will be strings
	myMap := make(map[string]int)

	// This adds a key of "value1" and assigns it a value of 25
	myMap["value1"] = 25

	fmt.Println(myMap)

	// This adds a second value
	myMap["value2"] = 26

	// This print will show both key value pairs
	fmt.Println(myMap)

	// We can also target specifc values like this
	fmt.Println("This is the value of key value2: ", myMap["value2"])

	delete(myMap, "value1")

	// This print should only include "value2"
	fmt.Println(myMap)
}

// =======================
// NOTES
// =======================
// Maps are built on hash tables.

// the value of an uninitialized map is nil.
// // This is because maps are reference types.

// Again we can use make to create maps.
// //  We can also use the short hand.
