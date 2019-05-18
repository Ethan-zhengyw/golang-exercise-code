package main

import "fmt"

func findLongestNotRepeatedSubString(s string) int {
    lastOccured := make(map[byte]int)
    start, maxLength := 0, 0
    for i, ch := range []byte(s) {
        if lastI, ok := lastOccured[ch] ; ok && (lastI >= start) {
            start = lastI + 1
        }
        if i - start + 1 > maxLength {
            maxLength = i - start + 1
        }
        lastOccured[ch] = i
    }
    return maxLength
}

func main() {
    s1 := "abcabcbb"  // --> abc
    s2 := "bbbb"  // --> b
    s3 := "aabcdeefghijjke"  // --> efghij
    s4 := ""  // --> ""
    s5 := "1234567890abcde"  // --> "1234567890abcde"

    fmt.Println(s1, s2, s3)

    fmt.Println(
        findLongestNotRepeatedSubString(s1),
        findLongestNotRepeatedSubString(s2),
        findLongestNotRepeatedSubString(s3),
        findLongestNotRepeatedSubString(s4),
        findLongestNotRepeatedSubString(s5),
    )
}
