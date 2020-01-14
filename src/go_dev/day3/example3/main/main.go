package main

import "fmt"

type op_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operation(op op_func, a, b int) int {
	return op(a, b)
}

func concat(a string, arg ...string) string {
	result := a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return result
}

func deffer_test() {
	var i int = 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}(a, b)
	return result
}

func main() {
	//var n int
	//n = rand.Intn(100)
	//
	//for {
	//	var input int
	//	fmt.Scanf("%d", &input)
	//	flag := false
	//	switch {
	//	case input == n:
	//		fmt.Println("OK")
	//		flag = true
	//	case input > n:
	//		fmt.Println(">")
	//	case input < n:
	//		fmt.Println("<")
	//	}
	//	if flag {
	//		break
	//	}
	//}

	//str := "hello world,中国"
	//for index, val := range str {
	//	fmt.Printf("index[%d] val[%c] len(%d)\n", index, val, len([]byte(string(val))))
	//}
	//
	//i := 0
	//HEAR:
	//	fmt.Println(i)
	//	i++
	//	if i == 5 {
	//		return
	//	}
	//	goto HEAR

	fmt.Println(operation(add, 100, 399))

	fmt.Println(concat("a", "ccc", "dddd", "eeee"))

	deffer_test()

	fmt.Println(test(100, 200))
}
