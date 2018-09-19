package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	rdr := strings.NewReader(`{"First":"Lucas","Last":"Werner","Age":26}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println("First: ", p1.First)
	fmt.Println("Last: ", p1.Last)
	fmt.Println("Age: ", p1.Age)
}
