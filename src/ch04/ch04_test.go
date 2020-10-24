package ch04_test

import (
	. "ch04"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestDoGlobalScope(t *testing.T) {
	DoGlobalScope()
}

func TestDoTypelessConstant(t *testing.T) {
	DoTypelessConstant()
}

func TestCh04Random(t *testing.T) {
	for i := 0; i < 10; i++ {
		println("rand", i, rand.Int())
	}
	println()
	for i := 0; i < 5; i++ {
		println("rand.Intn", i, rand.Intn(8))
	}
	println()
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		fmt.Printf("rand.Float32() %d %2.2f\n", i, rand.Float32() * 100)
	}
}

func TestCh04StrFunc(t *testing.T)  {
	println(strings.HasPrefix("Hello", "He"))
	println(strings.Contains("Hello", "He"))
	println(strings.Index("Hello", "l"))
	println(strings.LastIndex("Hello", "l"))
	println(strings.ToUpper("Hello"))
	println(strings.TrimSpace("  Hello  "))
	println(strings.TrimLeft("  Hello  ", " "))
	println(strings.TrimRight("  Hello  ", " "))
	println(strings.Fields("Hello world")[0])
	println(strings.Split("Hello world", " ")[0])
}

func shouldReturn() int {
	if true {
		return 1
	} else {
		return 2
	}
}

func TestCh04ForceReturn(t *testing.T)  {
	shouldReturn()
}

func TestCh04SPrintf(t *testing.T)  {
	a := "hello %s"
	a = fmt.Sprintf(a, runtime.GOOS)
	println("SPrintf", a)
}
