package main

import "fmt"

func newCounter() func() int {
	count := 0 // free variable, persisted data
	return func() int {
		count++
		return count
	}
}

func main() {
	count := 0

	counter := func() int {
		count++
		return count
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", counter())
	}
	fmt.Println(count)

	counter = newCounter()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", counter())
	}
	fmt.Println(count)
}
