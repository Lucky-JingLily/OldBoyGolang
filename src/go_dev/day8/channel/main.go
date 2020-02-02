package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 10)
	intChan2 := make(chan int, 10)

	go func() {
		var i int
		for {
			intChan <- i
			time.Sleep(time.Second)
			intChan2 <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()

	// 关闭channel
	//close(intChan)

	//for {
	//	var b int
	//	b, ok := <- intChan
	//	if !ok {
	//		fmt.Println("chan is close")
	//		break
	//	}
	//	fmt.Println(b)
	//}

	//for v := range intChan {
	//	fmt.Println(v)
	//}

	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
		case v := <-intChan2:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("time out")
			time.Sleep(time.Second)
		}
	}

}
