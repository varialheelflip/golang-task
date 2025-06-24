package main

import "time"

func main() {
	tasks := []func(){print1, print2, print3}
	taskSchedule(tasks)
	time.Sleep(1 * time.Second)
}
