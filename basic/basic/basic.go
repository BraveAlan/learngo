package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 作用域在包内部
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// bb := true //这是不可以的

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) // "%q"代表quatation
}

func variabelInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTyeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	var s = "abc"
	fmt.Println(a, b, c, d, s)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4 // 没定类型的时候，即可以做int，也可以做float，这个相当于是“文本替换”的动作
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello Wrold")
	variableZeroValue()
	variabelInitValue()
	variableTyeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	enums()
}
