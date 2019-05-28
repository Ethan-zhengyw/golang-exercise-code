package main

import "fmt"

func doWork1(a, b int) int {
	fmt.Printf("doWork1 called %d + %d = %d\n", a, b, a+b)
	return a + b
}

func doWork2(a int) int {
	fmt.Printf("doWork2 called 2 * %d = %d\n", a, 2*a)
	return a * 2
}

func doWork3(a int) int {
	fmt.Printf("doWork3 called 3 * %d = %d\n", a, 3*a)
	return a * 3
}

func main() {

	a, b := 0, 1

	go func() {
		result := doWork1(a, b)
		result = doWork2(result)
		result = doWork3(result)
		// Use the final result
	}()
	fmt.Println("hi!")
}
