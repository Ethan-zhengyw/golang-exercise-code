package main

import "fmt"
import "learngo/superqueue"

func main() {
    q := superqueue.Queue{}
    q.Push(1)
    q.Push("abc")
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
