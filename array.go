package main

import "fmt"

func array() {
    var arr1 [3]int
    arr2 := [3]int{1, 2, 3}
    arr3 := [...]int{1, 2, 3, 4}

    var grid [4][5]int
    
    fmt.Println(arr1, arr2, arr3)
    fmt.Println(grid)
}

func itera() {
    arr := [...]int{2, 4, 6, 8, 10}
    for i := range arr {
        fmt.Println(i)
    }

    for i, v := range arr {
        fmt.Println(i, v)
    }

    for _, v := range arr {
        fmt.Println(v)
    }
}

func modify(arr *[3]int) {
    (*arr)[0] = 100
}

func modify2(arr *[3]int) {
    arr[0] = 200
}

func main() {
    array()
    itera()
    
    var arr1 [3]int
    arr2 := [3]int{1, 2, 3}
    arr3 := [...]int{1, 2, 3, 4}

    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr3)
    modify(&arr2)
    fmt.Println(arr2)
    modify2(&arr2)
    fmt.Println(arr2)
}
