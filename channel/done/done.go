package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {

	for n := range w.in { // 会检测close
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

// 建了一个channel
// 开一个goroutine
// 立刻返回，收到这个channel的人，只能对它发数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	// 这个channel要做的事
	go doWorker(id, w)
	return w
}

func chanDemo() {
	//var c chan int // c == nil 不可以直接用
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

}

func main() {

	chanDemo()

}
