package main

import "fmt"

func main() {
	// Unlike arrays, slices do not have a defined length
	mySlice := []int{1, 2, 3, 4, 5}

	// This prints the whole slice
	fmt.Println(mySlice)

	// We can also create arrays with the "make" and "new functions"
	// In this example the 50 is the starting length of the slice,
	// and then 100 is the capacity for the underlying array.
	makeSlice := make([]int, 50, 100)

	// This will print 50 zeroes because that was our startin value
	fmt.Println(makeSlice)

	// This does the same thing as make just with different syntax
	newSlice := new([100]int)[0:50]
	fmt.Println(newSlice)

	// Here we will manipulate a slice and see how it responds to adding more values
	exampleSlice := make([]int, 0, 5)

	for i := 0; i < 10; i++ {
		exampleSlice = append(exampleSlice, i)
		fmt.Println("Length: ", len(exampleSlice), " Capacity: ", cap(exampleSlice), " Value: ", exampleSlice[i])
	}

	// Notice that the capacity doubles when our length exceeds our initial capacity
	// append is very similar to push() in JavaScript

}
