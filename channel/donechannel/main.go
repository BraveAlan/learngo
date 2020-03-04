package main

import "time"
import "fmt"

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)
	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case v, ok := <-done:
				fmt.Println("child process interrupt...", v, ok)
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
