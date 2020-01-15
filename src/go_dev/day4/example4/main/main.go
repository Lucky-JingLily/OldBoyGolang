package main

import (
	"fmt"
	"sort"
)

func testMap() {
	var a map[string]string
	a = make(map[string]string, 5)
	a["add"] = "+"
	fmt.Println(a)

	var b map[string]map[string]string = make(map[string]map[string]string, 10)
	b["key"] = make(map[string]string, 1)
	b["key"]["dfa"] = "sdfa"
	fmt.Println(b)

	val, ok := a["add"]
	if ok {
		fmt.Println(val)
	}
}

func testMapSort() {
	var a = make(map[int]int, 5)
	a[4] = 4
	a[7] = 4
	a[5] = 4
	a[3] = 4
	a[1] = 4

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}
func main() {
	//f := adder()
	//fmt.Println(f(1))
	//fmt.Println(f(100))
	//fmt.Println(f(1000))
	testMapSort()
}
