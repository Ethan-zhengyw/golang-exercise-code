package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()

	fmt.Println(
		f(), // 1
		f(), // 1
		f(), // 2
		f(), // 3
		f(), // 5
	)
}
