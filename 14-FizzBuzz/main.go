package main

import "fmt"

func main() {
	// This exercise will check if a number is divisible by 2, 3 or both.
	// If divisible by 2 it will print Fizz.
	// If divisible by 3 it will print Buzz.
	// And if divisible by 2 and 3 it will print FizzBuzz.
	for i := 1; i <= 100; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%2 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	var userNum int
	fmt.Print("Please enter a number: ")

	fmt.Scan(&userNum)

	//This will do the same as the previous but will take a user input
	if userNum%2 == 0 && userNum%3 == 0 {
		fmt.Println("FizzBuzz")
	} else if userNum%2 == 0 {
		fmt.Println("Fizz")
	} else if userNum%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(userNum)
	}
}
