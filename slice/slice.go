package main

import "fmt"

func modify(arr []int) {
    arr[0] = 999
}

func printSlice(s []int) {
    fmt.Printf("%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}

func main() {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6}
    s := arr[2:6]
    fmt.Println("arr[2:6] is ", s)
    fmt.Println("arr[:6] is ", arr[:6])
    fmt.Println("arr[2:] is ", arr[2:])
    fmt.Println("arr[:] is ", arr[:])

    modify(s)
    fmt.Println(s)
    fmt.Println(arr)

    arr = [...]int{0, 1, 2, 3, 4, 5, 6}
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))
    s1 := arr[1:5]
    s2 := s1[2:6]
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
    fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

    s1 = append(s1, 10)
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))

    s1 = append(s1, 11)
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))

    s1 = append(s1, 12)
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))

    s1 = append(s1, 13)
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))

    fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

    var pureSlice []int
    fmt.Println(pureSlice)   
    for i := 0; i < 10; i++ {
        printSlice(pureSlice)
        pureSlice = append(pureSlice, 2 * i + 1)
    }
    fmt.Println(pureSlice)   

    fmt.Println("Creating slice...")
    ps1 := make([]int, 10)
    ps2 := make([]int, 13, 100)
    printSlice(ps1)
    printSlice(ps2)

    fmt.Println("Copying slice...")
    copy(ps2, s2)
    printSlice(ps2)

    fmt.Println("Deleting element...")
    ps2 = append(ps2[:2], ps2[3:]...)
    printSlice(ps2)

    front := ps2[0]
    fmt.Println("Poping from front...", front)
    ps2 = ps2[1:]
    printSlice(ps2)

    back := ps2[len(ps2) - 1]
    fmt.Println("Poping from back...", back)
    ps2 = ps2[:len(ps2) - 1]
    printSlice(ps2)
}
