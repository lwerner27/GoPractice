package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	// Basic Ascending Sort
	studyGroup := people{"Josh", "Lucas", "Abbie", "Zak"}
	// Sort Strings will work because our people type
	// satisfies the slice of Interface interface for sort
	fmt.Println("before sort: ", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("after sort: ", studyGroup)
	// Just another example of using the sort package
	numbers := []int{1, 5, 7, 2, 9, 4, 6, 3, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)

}
