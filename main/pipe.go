package main

import "fmt"

func test_pipe() {
	// chan管道，相当于队列
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	fmt.Println(len(pipe))
}
