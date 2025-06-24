package main

import (
	"fmt"
	"time"
)

func taskSchedule(tasks []func()) {
	for index, task := range tasks {
		index := index
		task := task
		go func() {
			start := time.Now().UnixMilli()
			task()
			end := time.Now().UnixMilli()
			fmt.Println("任务", index, "执行时间:", end-start, "毫秒")
		}()
	}
}

func print1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func print2() {
	for i := 11; i <= 20; i++ {
		fmt.Println(i)
	}
}

func print3() {
	for i := 21; i <= 30; i++ {
		fmt.Println(i)
	}
}
