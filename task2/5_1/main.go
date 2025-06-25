package main

import (
	"fmt"
	"sync"
	"time"
)

func plusOneThousand(count *int) {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 1000; i++ {
		*count++
	}
}

var mu sync.Mutex

func main() {
	input := 0
	for i := 0; i < 10; i++ {
		go plusOneThousand(&input)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(input)
}
