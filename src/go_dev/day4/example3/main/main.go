package main

import (
	"fmt"
	"sort"
)

func adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	b := a[2:4]
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = b[0:1]
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}

func testSliceCopy() {
	a := []int{1, 2, 3, 4, 5}

	b := make([]int, 10)
	copy(b, a)
	fmt.Println(b)

	c := make([]int, 5)
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3)
	c = append(c, 4)
	c = append(c, 5)
	fmt.Println(c)
}

func testSort() {
	var a = [...]int{1, 2, 4342, 42, 4342, 423}
	sort.Ints(a[:])
	fmt.Println(a)
	fmt.Println(sort.SearchInts(a[:], 2))
}
func main() {
	//f := adder()
	//fmt.Println(f(1))
	//fmt.Println(f(100))
	//fmt.Println(f(1000))
	testSort()
}
