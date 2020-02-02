package main

import (
	"fmt"
	"time"
)

func test() {
	var m map[string]int

	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
			return
		}
	}()

	m["haha"] = 11
}

func main() {
	for i := 0; i < 10; i++ {
		go test()
	}
	time.Sleep(time.Second * 100)
}
