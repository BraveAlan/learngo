package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) { // 非抢占式
			fmt.Println("hello from", i) // I/O goroutine切换点
		}(i)
	}
	time.Sleep(time.Millisecond)
}
