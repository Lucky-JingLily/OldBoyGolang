package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//var lock sync.MutexRW
var rwlock sync.RWMutex

func testRWLock() {
	var a = make(map[int]int, 5)
	a[4] = 4
	a[7] = 4
	a[5] = 4
	a[3] = 4
	a[1] = 4
	var count int32
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[4] = rand.Intn(100)
			rwlock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			rwlock.RLock()
			fmt.Println(b)
			rwlock.RUnlock()
			atomic.AddInt32(&count, 1)
		}(a)
	}

	time.Sleep(5000 * time.Millisecond)
	fmt.Println(atomic.LoadInt32(&count))
}

//func testLock() {
//	var a = make(map[int]int, 5)
//	a[4] = 4
//	a[7] = 4
//	a[5] = 4
//	a[3] = 4
//	a[1] = 4
//
//	for i := 0; i < 2; i++ {
//		go func(b map[int]int) {
//			lock.Lock()
//			b[4] = rand.Intn(100)
//			lock.Unlock()
//		}(a)
//	}
//
//	lock.Lock()
//	fmt.Println(a)
//	lock.Unlock()
//}

func main() {
	//f := adder()
	//fmt.Println(f(1))
	//fmt.Println(f(100))
	//fmt.Println(f(1000))
	testRWLock()
}
