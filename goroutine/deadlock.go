package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // 每一个goroutine都固定一个i
			for {
				a[i]++
				runtime.Gosched() // 手动交出控制权，不然会死锁
			}
		}(i)

		// 这里传进去的 i 会变成10
		// 报错 index out of range
		//go func() { // race conditon, 在命令行中执行 go run -race deadlock.go 检测数据访问的冲突
		//	fmt.Println("i = ", i)
		//	for {
		//		a[i]++            // 这里引用的i和外面for循环引用的i是相同的，当外面for循环执行完毕后，i变为10
		//		runtime.Gosched() // 手动交出控制权，不然会死锁
		//	}
		//}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(a)
}
