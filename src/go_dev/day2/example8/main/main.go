package main

import "fmt"

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

var c string

func main() {
	a := 5
	b := make(chan int, 1)

	var i int8 = 1
	var j int16 = (int16(i))

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	first := 100
	second := 200
	swap(&first, &second)
	fmt.Println("first=", first, "second=", second)

	c = "G"
	println(c)
	f1()
}

func f1() {
	c := "O"
	fmt.Println(c)
	f2()
}

func f2() {
	fmt.Println(c)
}
