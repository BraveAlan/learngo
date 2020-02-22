package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱中国!" // UTF-8 可变长，英文1字节，中文3字节

	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) // 解第一个rune
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 另外开了一个rune数组，
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

}
