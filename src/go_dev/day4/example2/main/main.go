package main

import (
	"errors"
	"fmt"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//b := 0
	//fmt.Println(100/b)
	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	test()
	var a []int
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	a = append(a, a...)
	fmt.Println(a)
}
