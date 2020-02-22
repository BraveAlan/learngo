package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) { // range s是不行的，这样遍历出来，他把中文还是算3个长的
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabc"),
		lengthOfNonRepeatingSubStr("aaa"),
		lengthOfNonRepeatingSubStr("a"),
		lengthOfNonRepeatingSubStr(""),
		lengthOfNonRepeatingSubStr("一二三二一"),
		lengthOfNonRepeatingSubStr("你们好"),
		)
}
