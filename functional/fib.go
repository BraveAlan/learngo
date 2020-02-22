package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1 1 2 3 5 8 13
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 函数是一等公民，给函数实现接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000000000000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { // 没有起始条件，也没有递增条件，只有一个结束条件，相当于while
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)

	//p := make([]byte, 2)
	//n, err := f.Read(p)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("read: ", n)
	//str := string(p)
	//fmt.Println(str)
	//fmt.Printf("%v, ", p)

}
