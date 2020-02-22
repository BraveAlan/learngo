package main

import (
	"bufio"
	"fmt"
	"os"

	"imooc.com/sb/learngo/functional/fib"
)

func tryDefer() {
	//defer fmt.Println(1) // 函数退出后执行
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occcured")

	n := 100
	fmt.Println("n1=", n)
	defer fmt.Println("defer n=", n) // 参数在defer语句时计算
	n += 100
	fmt.Println("n2=", n)

}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666) // r4w2x1
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}
