package main

import (
	"fmt"
)

func main() {

	myWords := make([]string, 0, 10)
	myWords = append(myWords, "Lucas")
	myWords = append(myWords, "James")
	myWords = append(myWords, "Werner")
	myWords = append(myWords, "Amy")
	myWords = append(myWords, "Jo")
	myWords = append(myWords, "Hujet")

	for i := range myWords {
		fmt.Println(myWords[i])
	}

}

// ============================================
// Notes about Hash Tables
// ============================================
// * Hash tables are really fast at for searches
// * Hash tables use the idea of "buckets"
// // - We want even distribution of our items across our buckets
// // - Buckets will be determined by an algorithm assigning a
// // unique value to the items we are trying to store (Hash values)
