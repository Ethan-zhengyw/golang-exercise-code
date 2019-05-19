package main

import "fmt"
import "strings"

func main() {

	s0 := "a b 	c	d eeef"
	s1 := "aaa"
	s2 := "bBb"
	s3 := "cCc"
	sArr := []string{"java", "python", "go"}

	fmt.Println("s0 is:", s0)
	fmt.Println("s1 is:", s1)
	fmt.Println("s2 is:", s2)
	fmt.Println("s3 is:", s3)

	fields := strings.Fields(s0)
	fmt.Println("strings.Fields(s0) is: ", fields)

	splits := strings.Split(s0, "ee")
	fmt.Println("strings.Split(s0, \"ee\") is: ", splits)

	joined := strings.Join(sArr, "+")
	fmt.Println("strings.Join(sArr, "+") is: ", joined)

	fmt.Println("strings.Contains(s0, \"d e\")", strings.Contains(s0, "d e"))
	fmt.Println("strings.Contains(s0, \"d   e\")", strings.Contains(s0, "d  e"))

	fmt.Println("strings.Index(s0, \"d e\")", strings.Index(s0, "d e"))
	fmt.Println("strings.Index(s0, \"d   e\")", strings.Index(s0, "d  e"))

	fmt.Println("strings.ToLower(s2) is:", strings.ToLower(s2))
	fmt.Println("strings.ToUpper(s2) is:", strings.ToUpper(s2))

	fmt.Println("strings.Trim(s3, \"c\") is:", strings.Trim(s3, "c"))
	fmt.Println("s3 is:", s3)
	fmt.Println("strings.TrimLeft(s3, \"c\") is:", strings.TrimLeft(s3, "c"))
	fmt.Println("s3 is:", s3)
	fmt.Println("strings.TrimRight(s3, \"c\") is:", strings.TrimRight(s3, "c"))
	fmt.Println("s3 is:", s3)
}