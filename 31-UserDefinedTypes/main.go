package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	// These are examples of struct literals
	p1 := person{"Lucas", "Werner", 26}
	p2 := person{"Amy", "Werner", 25}

	fmt.Println("First Name: ", p1.first, " Last Name: ", p1.last, " Age: ", p1.age)
	fmt.Println("First Name: ", p2.first, " Last Name: ", p2.last, " Age: ", p2.age)
}
