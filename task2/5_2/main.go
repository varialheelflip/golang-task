package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func plusOneThousand(count *int64) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(count, 1)
	}
}

func main() {
	var input int64
	for i := 0; i < 10; i++ {
		go plusOneThousand(&input)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(input)
}
