package main

import "fmt"
import "learngo/queue"

func main() {
    q := queue.Queue{}
    q.Push(1)
    q.Push(2)
    q.Push(3)

    fmt.Println(q)
    fmt.Println(q.IsEmpty())

    q.Pop()
    fmt.Println(q)
    fmt.Println(q.IsEmpty())

    q.Pop()
    fmt.Println(q)
    fmt.Println(q.IsEmpty())

    q.Pop()
    fmt.Println(q)
    fmt.Println(q.IsEmpty())
}
