package main

import "fmt"

// regular function definition(with name)
func namedFunc() {
	fmt.Println("I am a named function.")
}

func main() {

	namedFunc()

	var anonymousFunc func()

	anonymousFunc = func() {
		fmt.Println("I am a anonymous function.")
	}
	anonymousFunc()

	// update function content while program running
	anonymousFunc = func() {
		fmt.Println("I am a updated anonymous function.")
	}
	anonymousFunc()

	// set anonymous function with regular function
	anonymousFunc = namedFunc
	fmt.Println("set anonymous function with regular function.")
	anonymousFunc()

}
