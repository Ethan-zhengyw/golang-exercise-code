package main

import "fmt"
import "unicode/utf8"

func main() {
	s := "hello爸爸妈妈！good."

	fmt.Println(s, len(s))
	fmt.Printf("%s\n", s)	
	fmt.Printf("%X\n", s)

	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune Count in String s: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("(%c, %d) ", ch, size)
		bytes = bytes[size:]
	}
	fmt.Println()

	sRune := []rune(s)
	fmt.Println("length of sRune: ", len(sRune))
	for i, ch := range sRune {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()
}
