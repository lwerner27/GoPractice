package main

import (
	"fmt"
)

func main() {
	// =======================================================
	// Most Basic for loop example
	// Prints nummbers 0 through 9
	// Basically the exact same syntax as javascript
	// =======================================================

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// =======================================================
	// Example of Nested Loop
	// =======================================================

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("This is the outer loop: ", i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Println("This is the INNER loop: ", j)
	// 	}
	// }

	// =======================================================
	// Example of Conditional Loop
	// =======================================================

	// i := false
	// j := 0

	// for i == false {
	// 	// Print the value of j
	// 	fmt.Println(i)
	// 	fmt.Println(j)

	// 	// If j is equal to 9 change i to true
	// 	if j == 9 {
	// 		i = true
	// 	}

	// 	// Increment j if not equal to 9
	// 	j++
	// }

	// fmt.Println(i)
	// fmt.Println("The Loop has ended")

	// =======================================================
	// Using breaks and continues
	// =======================================================
	i := 0
	for {
		// This i is divisible by 2 continue the
		// This will skip the code that follows this block
		if i%2 == 0 {
			i++
			continue
		}

		fmt.Println(i)
		// If i equals 9 end the loop
		if i == 9 {
			break
		}

		i++

	}

	fmt.Println("The loop has ended.")
}
