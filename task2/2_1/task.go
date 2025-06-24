package main

import "fmt"

func printNum() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("goroutine1: ", i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("goroutine2: ", i)
			}
		}
	}()
}
