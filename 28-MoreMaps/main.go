package main

import (
	"fmt"
)

// This is another way to create a map
// Note the {} after the second "string"
// Without these you will get a nil on your map
// You will not be able to add values to it
var mySalutations = map[string]string{}

// One more way to create maps
// this example gives us some startin values
var myMenu = map[string]string{
	"Tacos":      "A delicious Mexican dish.",
	"Hamburgers": "An American classic.",
}

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Luke"] = "Good morning."
	myGreeting["Tim"] = "Good afternoon."

	mySalutations["Luke"] = "Goodbye."
	mySalutations["Tim"] = "So long."

	fmt.Println(myGreeting)
	fmt.Println(mySalutations)
	fmt.Println(myMenu)
}
