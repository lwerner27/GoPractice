package main

import (
	"fmt"
)

// ========================
// Original zero function
// Not passing mem address
// ========================
// func zero(x int) {
// 	x = 0
// }

// ===============================
// OLD NON WORKING ZERO FUNCION
// ================================
// func main() {
// 	x := 5
// 	// this will change the value
// 	// But does not change the value stored in memory
// 	zero(x)
// 	fmt.Println(x) // Should still be 5
// }

// Unlike the previous 0 function this function expects a memory address not a value
// It then alters the value stored at the mem address as opposed to the value stored at the new mem address
func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	fmt.Println("This is the original value of x: ", x)
	zero(&x)
	fmt.Println("This is the new value of x: ", x)
}
