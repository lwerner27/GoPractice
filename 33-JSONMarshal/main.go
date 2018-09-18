package main

import (
	"encoding/json"
	"fmt"
)

// Person is a custom type
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Lucas", "Werner", 26, 007}
	bs, _ := json.Marshal(p1)
	fmt.Printf("%T \n", bs)
	// You shouldn't see notExported in the strigified version of your struct
	// This is because we chose not to export we not this because the idetifier is capitalized
	fmt.Println(string(bs))

}

// Marshal
// ** Marshal is similar to JSON to stringify
// Converts JSON to a string
