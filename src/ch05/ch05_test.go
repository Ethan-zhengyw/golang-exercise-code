package ch05_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCh05CommaOKPattern(t *testing.T)  {
	res, ok := strconv.Atoi("1234")
	println(res, ok)
}

func TestCh05CommaNotOKPattern(t *testing.T) {
	res, err := strconv.Atoi("123a4")
	if err != nil {
		println("exit with error: %v!", err)
		return
	}
	println(res, err)
}


func atoi (s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

func TestCh05WrapAtoiFunc(t *testing.T)  {
	println(atoi("hello")); println(atoi("1234"))

	count, err := fmt.Println(1)
	println(count, err)
}

func TestCh05SwitchFallthrough(t *testing.T)  {
	a := 0
	b := 1
	switch a {
	case 0:
	case 1:
		println("TestCh05SwitchFallthrough", 1)
	case 2:
		println("TestCh05SwitchFallthrough", 2, b)
	}

	switch b, a := 1, 0; {
	case a == 0:
		println("==0", b)
	}
}

func TestCh05SwitchFallthrouth2(t *testing.T)  {
	k := 6
	switch k {
	case 4: fmt.Println("was <= 4"); fallthrough;
	case 5: fmt.Println("was <= 5"); fallthrough;
	case 6: fmt.Println("was <= 6"); fallthrough;
	case 7: fmt.Println("was <= 7"); fallthrough;
	case 8: fmt.Println("was <= 8"); fallthrough;
	case 2: fmt.Println("was <= 2"); fallthrough;
	default: fmt.Println("default case")
	}
}

func TestCh05For(t *testing.T)  {
	for i := 0; i < 5; i++ {
		println("for loop now", i)
	}
}

func TestCh05ForString(t *testing.T)  {
	str1 := "世界你好！"
	fmt.Printf("length of %s is %d\n", str1, len(str1))

	for i := 0; i < len(str1); i++ {
		fmt.Printf("the %d character of str %s is: %c\n", i, str1, str1[i])
	}
}

func TestCh05GoTo(t *testing.T) {
	i := 0
	loop: if i < 15 {
		println("loop ", i)
		i++
		goto loop
	}
}

func TestCh05FizzBuzz(t *testing.T)  {
	for i := 0; i < 100; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			println("FizzBuzz")
		case i % 3 == 0:
			println("Fizz")
		case i % 5 == 0:
			println("Buzz")
		default:
			println(i)
		}
	}
}

func TestCh05ForRange(t *testing.T)  {
	str1 := "世界你好！"
	for i, c := range str1 {
		fmt.Printf("The %d character of str %s is %c\n", i, str1, c)
	}
}
