package main

import "fmt"
import "math"
import "math/cmplx"

var (
    aa = 3
    bb = "111"
    cc = true
)

var dd, ee, ff = 4, "444", false

func variableZeroValue() {
    var a int
    var b string
    fmt.Println(a, b)
    fmt.Printf("%d %s\n", a, b)
    fmt.Printf("%d %q\n", a, b)
}

func variableInitializeValue() {
    var a, b int = 2, 4
    var c string = "abc"
    fmt.Println(a, b, c)
}

func variableTypeDeduction() {
    var a, b, c, d = 2, 4, "abc", false
    fmt.Println(a, b, c, d)
}

func complex() {
    c := 3 + 4i
    fmt.Println(cmplx.Abs(c))
}

func euler() {
    fmt.Println(cmplx.Pow(
        math.E, 1i * math.Pi) + 1)
}

func euler2() {
    fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
    var a, b int = 3, 4
    var c int

    c = int(math.Sqrt(float64(a * a + b * b)))
    fmt.Println(c)
}

func consts() {
    const filename = "file.txt"
    const a, b = 3, 4
    var c int
    c = int(math.Sqrt(a * a + b * b))
    fmt.Println(filename, c)
}

func enums() {
    const (
        cpp = iota
        _
        java
        python
        golang
    )
    fmt.Println(cpp, java, python, golang)

    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
    fmt.Println("Hello World!")
    variableZeroValue()
    variableInitializeValue()
    variableTypeDeduction()
    fmt.Println(aa, bb, cc)
    fmt.Println(dd, ee, ff)
    complex()
    euler()
    euler2()
    triangle()
    consts()
    enums()
}
