package ch09_test

import (
	"container/list"
	"fmt"
	"regexp"
	"sync"
	"testing"
	"unsafe"
)

func TestDoubleLink(t *testing.T) {
	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v\n", item.Value)
	}

}

func TestUnsafe(t *testing.T) {
	var a *int
	println(unsafe.Sizeof(a))
	println(unsafe.Alignof(a))

	b := 1234
	println(unsafe.Sizeof(b))
	println(unsafe.Alignof(b))

	// println(unsafe.Offsetof(a))
}

func TestRegExp(t *testing.T) {
	s := "hello 1.1, world 2.2, ! 3.3"
	p := "[0-9]+.[0-9]+"

	res, e := regexp.Match(p, []byte(s))
	println(res, e)

	re, _ := regexp.Compile(p)
	s2 := re.ReplaceAllString(s, "##.##")
	println(s2)
}

func TestMutex(t *testing.T) {
	mu := sync.Mutex{}
	rwMu := sync.RWMutex{}

	mu.Lock()
	defer mu.Unlock()
	println("hello mu lock")
	rwMu.RLock()
	defer rwMu.RUnlock()
	println("hello rwMu lock")
}

func TestOnceDoFunc(t *testing.T) {
	f := func() {println("hello i should show only once.")}
	once := sync.Once{}
	once.Do(f)
	once.Do(f)
	once.Do(f)
}
