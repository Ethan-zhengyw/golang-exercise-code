package main

import "fmt"

func swap(a, b int) (int, int) {
    return b, a
}

func swap_by_p(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    a, b := 3, 4

    a, b = swap(a, b)
    fmt.Println(a, b)

    swap_by_p(&a, &b)
    fmt.Println(a, b)
}
