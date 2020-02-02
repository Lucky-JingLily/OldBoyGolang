package main

import (
	"fmt"
)

func calc(taskChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resultChan <- v
		}
	}

	fmt.Println("exit")
	exitChan <- true
}

func send(intChan chan int, exitChan chan bool) {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	exitChan <- true
	fmt.Println("send exit")
}

func recive(intChan chan int, exitChan chan bool) {
	for i := range intChan {
		fmt.Println(i)
	}
	exitChan <- true
	fmt.Println("recive exit")
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	//go func() {
	//	for v := range resultChan {
	//		fmt.Println(v)
	//	}
	//}()

	// 等待8个goroutine退出，<-不关注exitChan里面的值
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Println("exit,", i, "goroutine")
		}
		close(resultChan)
	}()

	for v := range resultChan {
		//fmt.Println(v)
		_ = v
	}
	//time.Sleep(time.Second * 10)

	intChan = make(chan int, 10)
	exitChan = make(chan bool, 2)

	go send(intChan, exitChan)
	go recive(intChan, exitChan)

	var tool int = 0
	for i := range exitChan {
		fmt.Println(i)
		tool++
		if tool == 2 {
			break
		}
	}
}
