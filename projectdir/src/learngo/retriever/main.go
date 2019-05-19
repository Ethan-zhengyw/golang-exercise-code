package main

import "fmt"
import "time"
import "learngo/retriever/mock"
import "learngo/retriever/real"

type Retriever interface {
    Get(url string) string
}

func download(r Retriever, url string) string {
    return r.Get(url)
}

func inspect(r Retriever) {
    fmt.Printf("%T %v\n", r, r)
    switch t := r.(type) {
        case mock.Retriever:
            fmt.Println(t.Contents)       
        case *real.Retriever:
            fmt.Println(t.UserAgent)       
            fmt.Println(t.TimeOut)       
    }
}

func inspectByAssertion(r Retriever) {
    fmt.Printf("%T %v\n", r, r)

    if mockRetriever, ok := r.(mock.Retriever); ok {
        fmt.Println(mockRetriever.Contents)
    } else {
        fmt.Println("not a mock.Retriever")
    }

    if realRetriever, err := r.(*real.Retriever); err {
        fmt.Println(realRetriever.UserAgent)
        fmt.Println(realRetriever.TimeOut)
    } else {
        fmt.Println("not a real.Retriever")
    }
}

func main() {
    r := mock.Retriever{Contents: "This is a fake www.baidu.com!"}
    contents := download(r, "www.baidu.com")
    inspect(r)
    inspectByAssertion(r)

    rReal := &real.Retriever{
        UserAgent: "mozilla/5.0",
        TimeOut: time.Minute,
    }
    fmt.Printf("%T %v\n", rReal, rReal)
    contents = download(rReal, "http://www.imooc.com")
    inspect(r)
    inspectByAssertion(rReal)
    fmt.Println(contents[:10])
}
