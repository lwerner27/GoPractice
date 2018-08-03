package main

import (
	"fmt"
)

// This is a constant
// Also multiple initialization
const (
	Pi       = 3.14
	Language = "Go"
)

// Iotas are verys small numbers
// They start at 0 and increment by 1
const (
	_  = iota             // we are throwing out iota 0
	KB = 1 << (iota * 10) // iota = 1
	MB = 1 << (iota * 10) // iota = 2
	GB = 1 << (iota * 10) // iota = 3
)

// This is another way to write a constant
const name string = "Lucas Werner"

func main() {
	fmt.Println(Pi)
	fmt.Println(Language)
	fmt.Println(name)
}
