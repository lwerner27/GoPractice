package main

import (
	"encoding/json"
	"fmt"
)

// Person type for use in this example
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person

	fmt.Println("Initial value of p1.First: ", p1.First)
	fmt.Println("Initial value of p1.Last: ", p1.Last)
	fmt.Println("Initial value of p1.Age: ", p1.Age)

	bs := []byte(`{"First": "Lucas", "Last": "Werner", "Age": 26}`)
	// Here we are going to unmarshal the slice of bytes
	// we just made that represents the JSON object above

	// Notice how we pass in the addres of our p1 variable
	// this is so we can change the values that are being stored
	json.Unmarshal(bs, &p1)

	fmt.Println("==========================================")
	fmt.Println("New value of p1.First: ", p1.First)
	fmt.Println("New value of p1.Last: ", p1.Last)
	fmt.Println("new value of p1.Age: ", p1.Age)

}
