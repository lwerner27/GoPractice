package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i <= 1000; i++ {
		fmt.Println("Foo: ", i)
		// This line will just make it a lot easier to see that
		// The two functions are running concurrently
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 1000; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	// wg.Add(2) and wg.Wait() will make our main function wait
	// For two other functions to report done before it ends itself
	wg.Add(2)
	// The go keyword will make it so that any
	// Functions that have the keyword will run concurrently
	go foo()
	go bar()
	wg.Wait()
}
