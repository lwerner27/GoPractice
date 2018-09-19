package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{"Lucas", "Werner", 26}
	// Here NewEncoder returns a point to an encoder
	// This is what allows us use the Encode method to encode our data
	json.NewEncoder(os.Stdout).Encode(p1)
}

// Encoding is sending out data
// Decoding is about reading data
