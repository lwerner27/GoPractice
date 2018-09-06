package main

import "fmt"

// Very similar to class objects in other languages
type person struct {
	first string
	last  string
	age   int
}

// This is a type that inherits the fields from person
type family struct {
	person
	quantity int
}

// Note that both family and person are lower case
// This means they are unexported
// This is similar to private objects in other languages

func main() {
	// We are creating a variable using our type of "person"
	person1 := person{"Lucas", "James", 26}

	// On the lines below we are accessing the different fields or our type
	fmt.Println("First Name: ", person1.first)
	fmt.Println("Last Name: ", person1.last)
	fmt.Println("Age: ", person1.age)

	// Here we can showcase the inheritance
	// We do this by passing in  a previously created type
	family1 := family{person1, 1}

	fmt.Println(family1)
}
