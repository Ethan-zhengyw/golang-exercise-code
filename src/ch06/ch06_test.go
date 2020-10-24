package ch06_test

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Pop(st string) uint8 {
	var v uint8
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			return v
		}
	}
	return v
}

func TestCh06MissingReturn(t *testing.T)  {
	println(Pop("hello world!"))
}

var num int = 10
var numx2, numx3 int


func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}


func TestCh06MultipleValue(t *testing.T) {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func ReceiveMultipleArgs(arg ...int)  {
	println(arg[0], arg[1], arg[2])
}

func TestCh06MultipleArgs(t *testing.T) {
	arr := []int{1, 2, 3}
	ReceiveMultipleArgs(arr...)
}

type Options struct {
	var1 int
	str1 string
}

func ReceiveOptions(options Options)  {
	fmt.Printf("'%d', '%s'\n", options.var1, options.str1)
}

func TestCh06SendOptions(t *testing.T) {
	ReceiveOptions(Options{var1:1})
	ReceiveOptions(Options{str1:"str1"})
	ReceiveOptions(Options{var1:1, str1:"str1"})
}

func ReceiveInterfaces(args ...interface{})  {
	for idx, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Printf("%d is int, a is %v\n", arg, idx)
		case string:
			fmt.Printf("%s is string, a is %v\n", arg, idx)
		default:
			fmt.Printf("%v type unknown, a is %v\n", arg, idx)
		}
	}
}

func TestCh06SendInterface(t *testing.T)  {
	ReceiveInterfaces(5, "hello", 1.32)
}

func DeferFunc1() {
	println("DeferFunc1 called")
}

func TestCh06DeferFunc(t *testing.T) {
	println("start TestCh06DeferFunc")
	defer DeferFunc1()
	println("end TestCh06DeferFunc")
}

func TestCh06DeferFuncOrder(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}

func Enter(str string) {
	println("enter", str)
}

func Leave(str string)  {
	println("leave", str)
}

func TestCh06Tracing(t *testing.T) {
	Enter("TestCh06Tracing")
	defer Leave("TestCh06Tracing")
	println("TestCh06Tracing executing")
}

func Enter2(str string) string {
	println("enter", str)
	return str
}

func TestCh06Tracing2(t *testing.T) {
	defer Leave(Enter2("TestCh06Tracing2"))
	println("TestCh06Tracing executing")
}

func DebugFunc(var1 string) (res int, err error) {
	defer func() {
		log.Printf("DebugFunc(%q) = (%d, %v)", var1, res, err)
	}()
	return 1, io.EOF
}

func TestCh06Debug(t *testing.T)  {
	DebugFunc("hello")
}

func TestCh06Cap(t *testing.T) {
	a := []int{1, 2, 3}
	println(len(a), cap(a))

	b := append(a, 4)
	println(len(a), cap(a))
	println(len(b), cap(b))

	b = append(b, 5)
	println(len(a), cap(a))
	println(len(b), cap(b))

	b = append(b, 6)
	println(len(a), cap(a))
	println(len(b), cap(b))

	b = append(b, 7)
	println(len(a), cap(a))
	println(len(b), cap(b))

	b = append(b[:3], b[4:]...)
	println(len(a), cap(a))
	println(len(b), cap(b))

	c := append(a, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	c = append(c, 13)
	println(len(a), cap(a))
	println(len(b), cap(b))
	println(len(c), cap(c))
}

func TestCh06StringsMap(t *testing.T) {
	println(strings.Map(func(r rune) rune {
		if int(r) > 255 {
			return '?'
		} else {
			return r
		}
	}, "hello, 世界!"))
}

func TestCh06Closure(t *testing.T) {
	res1 := func() int {
		res2 := 0
		for i := 0; i <= 1e2; i++ {
			res2 += i
		}
		return res2
	}()
	println(res1)
}

func TestCh06Lambda(t *testing.T) {
	fv := func(str string) {println(str)}
	fv("Hello World!")
}

func DeferReturn() (ret int)  {
	defer func() {ret *= 2}()
	return 20
}

func TestCh06ReturnDefer(t *testing.T) {
	println(DeferReturn())
}

func TestCh06FibonacciClosure(t *testing.T) {
	f := func() func() {
		res := 0
		return func() {
			res++
			println(res)
		}
	}()

	f()
	f()
	f()
	f()
}

//type Suffix struct {
//	name string
//}
//
//type SuffixEnum struct {
//	Png Suffix
//	Jpeg Suffix
//}
//
////var Png = Suffix{name:"Png"}
////var Jpeg = Suffix{name:"Jpeg"}

func AddSuffixFactoryFunc(suffix string) func(fileName string) string  {
	return func(fileName string) string {
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		}

		return fileName + suffix
	}
}

func TestCh06FactoryFunc(t *testing.T) {
	addPngSuffix := AddSuffixFactoryFunc(".png")
	addJpegSuffix := AddSuffixFactoryFunc(".jpeg")

	println(addPngSuffix("hello"))
	println(addJpegSuffix("hello.jpeg"))
}

func TestCh06DebugWithClosureFunc(t *testing.T) {
	where := func() {
		pc, file, line, ok := runtime.Caller(1)
		log.Println(pc, file, line, ok)
	}

	where()

	where()
	where()

	where()
	log.SetFlags(log.Llongfile)
	log.Println("hello")

	log.SetFlags(log.Ldate)
	log.Println("hello")

	log.SetFlags(log.Lmicroseconds)
	log.Println("hello")

	log.SetFlags(log.Lshortfile)
	log.Println("hello")

	log.SetFlags(log.LstdFlags)
	log.Println("hello")

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Println("hello")

	where2 := log.Println

	where2()

	where2()
}

func TestCh06TimingFunc(t *testing.T) {
	start := time.Now()
	time.Sleep(1e6)
	end := time.Now()
	delta := end.Sub(start)
	log.Printf("TestCh06TimingFunc takes %v", delta)
}

func Timing(start time.Time)  {
	fmt.Printf("func took %v", time.Now().Sub(start))
}

func Fibs(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}
	return Fibs(n - 1) + Fibs(n - 2)
}


const SIZE  = 101
var FIBS [SIZE]uint64

func OptimizedFibs(n uint64) uint64  {
	if n == 0 || n == 1 {
		if n == 1 {
			FIBS[1] = 1
		}
		return 1
	} else if FIBS[n] != 0 {
		return FIBS[n]
	} else {
		FIBS[n] = OptimizedFibs(n - 1) +  OptimizedFibs(n - 2)
		return FIBS[n]
	}
}

func TestCh06OptimizedFibs(t *testing.T) {
	defer Timing(time.Now())
	println(OptimizedFibs(40))
}

func TestCh06Fibs(t *testing.T) {
	defer Timing(time.Now())
	println(Fibs(40))
}
