package ch08_test

import (
	"fmt"
	"testing"
)

func TestMapCreateAndAssign(t *testing.T) {
	m1 := map[int]int{1: 1, 2:2}
	fmt.Printf("%v\n", m1)

	m2 := make(map[int]int)
	m2 = m1
	fmt.Printf("%v\n", m2)

	m2[1] = 10
	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", m2)
}

func TestMapInitialize(t *testing.T) {
	var map1 map[int]int = make(map[int]int)
	fmt.Printf("%v\n", map1)

	map2 := map[int]int{1:1, 2:2}
	fmt.Printf("%v\n", map2)
}

func TestMapByNew(t *testing.T) {
	mp := new(map[int]int)
	*mp = map[int]int{1: 1}
	fmt.Printf("%v\n", mp)
	(*mp)[1] = 2
	fmt.Printf("%v\n", mp)
}

func TestMapKeyNotPresented(t *testing.T) {
	map1 := map[int]int{1:1, 2:2}

	v1, ok1 := map1[1]
	println(v1, ok1)

	v2 := map1[1]
	println(v2)

	v3, ok3 := map1[3]
	println(v3, ok3)
}

func TestOmitArrayRange(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	for idx, val := range arr1 {
		println(idx, val)
	}

	for idx := range arr1 {
		println(idx)
	}

	for _, val := range arr1 {
		println(val)
	}
}

func TestMapDeleteKey(t *testing.T) {
	map1 := map[int]int{1:1, 2:2}
	fmt.Printf("%v\n", map1)
	delete(map1, 2)
	fmt.Printf("%v\n", map1)
}
