package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// This is an example of a point reciever
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	c := circle{5}
	// Here we need the "&" to pass in a pointer
	// as required by our area method
	info(&c)

}

// Value you recievers will also accept a pointer type
// This will not work the other way around though
