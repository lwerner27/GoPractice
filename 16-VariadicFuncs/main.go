package main

import (
	"fmt"
)

func main() {
	data := []float64{2, 3, 4, 5, 6}
	fmt.Println("The average is ", average(5, 4))

	// This will pull the date out of a list and send it to the function
	fmt.Println("The average is ", average(data...))
}

// This function has a variadic parameter
// This is signaled by the ... in the declartion
// This means when it calls it can take any number of arguements
// This fucntion also makes use of a slice which we will cover later
// But basically sf is an array of float64 numbers
func average(sf ...float64) float64 {
	total := 0.0

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
