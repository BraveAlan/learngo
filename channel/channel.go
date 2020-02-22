package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//for {
	//
	//	fmt.Printf("Worker %d received %d\n", id, <-c) // 如果发送方close，这样的写法会不停的接收零值
	//
	//	if n, ok := <-c; !ok { // 检测close
	//		break
	//	} else {
	//		fmt.Printf("Worker %d received %d\n", id, n)
	//	}
	//}

	// 另一种写法
	for n := range c { // 会检测close
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

// 建了一个channel
// 开一个goroutine
// 立刻返回，收到这个channel的人，只能对它发数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	// 这个channel要做的事
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int // c == nil 不可以直接用
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) // 建了缓冲区，goroutine不会频繁切换
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	// 发送方close
	c := make(chan int, 3) // 建了缓冲区，goroutine不会频繁切换
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	//chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
