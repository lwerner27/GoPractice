package main

import (
	"fmt"
)

// Square is a user defined type
// With an underlying type of struct
type Square struct {
	Side float64
}

// this is how we can attach a method to our Square type
func (z Square) area() float64 {
	return z.Side * z.Side
}

// Shape is another user defined type
// Only this time it has an underlying type of interface
// Now anything type that has the specified characteristics
// can be considered shape
// since our squre has an area method that takes no parameters and returns a float64
// it can be considered a shape and used where a shape can be used
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// now we can see that even tho our info function wants a element with a type of shape
// Giving info a square will work because it meets shapes criteria to be consider a shape
func main() {
	s := Square{10}
	info(s)
}

// Some people call this polymorphism and some people call this substitutabliity ??Spelling??
