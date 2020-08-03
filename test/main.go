package main

import (
	"fmt"
	"time"
)

var timeFormat = "2006-01-02 15:04:05"

func main() {
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	fmt.Printf("now: %s\n", now)
	testTime := now.Format(timeFormat)
	fmt.Printf("testTime: %v\n", testTime)
}
