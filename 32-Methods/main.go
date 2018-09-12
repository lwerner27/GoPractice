package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// This is how we can add methods to our new type
// the "(p person)" is a reciever and it tells this function what type to associate with
// similar to this in other languages (do not use this or self in Go)
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {

	// These are examples of struct literals
	p1 := person{"Lucas", "Werner", 26}
	p2 := person{"Amy", "Werner", 25}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

}
