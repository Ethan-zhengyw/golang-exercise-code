package ch07_test

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestCh07LoopArray(t *testing.T) {
	arr := [5]int{}

	for i := 0; i < len(arr); i++ {
		println(i, arr[i])
	}

	for i := range arr {
		println(i, arr[i])
	}

	arr2 := [...]string{"a", "b", "c"}

	for i := range arr2 {
		println(i, arr2[i])
	}
}

func TestCh07SliceVsArray(t *testing.T) {
	//s := []int
	a := [5]int{}
	log.Printf("%v", reflect.TypeOf(a))

	s := a[0:3]
	log.Printf("%v %d %d", reflect.TypeOf(s), len(s), cap(s))

	v := [...]int{}
	log.Printf("%v %d %d", reflect.TypeOf(v), len(v), cap(v))

	var ss []int
	log.Printf("%v", reflect.TypeOf(ss))

	sss := &a
	log.Printf("%v", reflect.TypeOf(sss))

	byMake := make([]int, 0)
	log.Printf("%v", reflect.TypeOf(byMake))

	s1 := []byte{'p', 'o', 'e', 'm'}
	println(s1, s1[0], s1[1], s1[2], s1[3], &s1[3])

	s2 := s1[2:]
	println(s2, s2[1])

	s2[1] = 't'
	println(s1, s1[0], s1[1], s1[2], s1[3])
	println(s2, s2[1])
}

func TestCh07BytesBuffer(t *testing.T) {
	var bb bytes.Buffer
	bb.WriteString("hello\nworld!")
	res := strings.Split(bb.String(), "\n")
	log.Printf("%q", res)

	//header, left := bb[:6], bb[7:]
	//header, left :=

	//var bbp *bytes.Buffer
	//bbp.WriteString("hello world!")
	//println(bbp.String())

}

func TestCh07DoubleSlice(t *testing.T) {

	arr := [...]int{10, 20, 30}

	for idx := range arr {
		arr[idx] *= 2
	}

	fmt.Printf("%v", arr)

}

func TestCh07Reslice(t *testing.T) {
	a := make([]int, 0, 10)

	fmt.Printf("len(a) = %d, cap(a) = %d\n", len(a), cap(a))

	for i := 0; i < 10; i++ {
		a := a[:i + 1]
		a[i] = i
		fmt.Printf("len(a) = %d, cap(a) = %d\n", len(a), cap(a))
	}
}

func TestCh07Copy(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := make([]int, 5, 10)
	copy(b, a[:])
	log.Printf("%v", b)

	b = append(b, 4, 5, 6)
	log.Printf("%v", b)
}

func TestCh07SortSlice(t *testing.T) {
	aa := [...]int{5, 3, 7, 8, 4}
	a := aa[:]
	println(sort.IntsAreSorted(a))
	println(a[sort.SearchInts(a, 7)])

	sort.Ints(a)
	println(sort.IntsAreSorted(a))
	println(a[sort.SearchInts(a, 7)])
}

func TestCh07Map(t *testing.T) {
	a := [...]int{1, 2, 3}
	mapFunc := func(f func(int) int, s []int) []int {
		for idx, item := range s {
			s[idx] = f(item)
		}
		return s
	}
	fmt.Printf("%v", mapFunc(func(i int) int { return 2 * i }, a[:]))
}

func TestMapSort(t *testing.T) {
	map1 := map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16,
		"indio": 87, "juliet": 65, "kilo": 43, "lima": 98}

	for key, val := range map1 {
		println(key, " ", val)
	}

	//keys := make([]string, len(map1))
	keys := make([]string, len(map1))
	idx := 0
	for k := range map1 {
		//keys = append(keys, k)
		keys[idx] = k
		idx++
	}

	sort.Strings(keys)
	for _, k := range keys {
		println(k, " ", map1[k])
	}
}
