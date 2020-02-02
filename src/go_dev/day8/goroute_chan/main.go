package main

import (
	"fmt"
	"time"
)

func write(intChan chan int) {
	for i := 0; i < 100; i++ {
		intChan <- i
	}
}

func read(intChan chan int) {
	for {
		a := <-intChan
		fmt.Println(a)
	}
}

func main() {
	intChan := make(chan int, 20)
	go write(intChan)
	go read(intChan)

	time.Sleep(time.Second * 10)
}
