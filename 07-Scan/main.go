package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64

	fmt.Print("Enter meters swam: ")

	// You need the & so your code knows where to store variable.
	fmt.Scan(&meters)

	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
