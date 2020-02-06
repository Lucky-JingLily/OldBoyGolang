package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 1)

	go func() {
		var i int
		for {
			// 写入，如果chan满的时候执行default
			select {
			case ch <- i:
			default:
				fmt.Println("chan full, timeout")
				time.Sleep(time.Second)
			}
			i++
		}
	}()

	for {
		v := <-ch
		fmt.Println(v)
	}
}
