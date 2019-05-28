package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 11, -5, 7, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)
	index := sort.SearchInts(numbers, 7)
	fmt.Println("7 is at index:", index)

	// reverse slice
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("Sorted:", numbers)

	// This closure accesses the numbers slice even though it is never passed in,
	// and returns true for any number that is greater than or equal to 7.
	index = sort.Search(len(numbers), func(i int) bool {
		return numbers[i] <= -6
	})
	fmt.Println("-6 is at index:", index)
}
