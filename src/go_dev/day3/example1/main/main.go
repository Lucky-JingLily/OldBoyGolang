package main

import (
	"fmt"
	"strconv"
)

func isPrim(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isShui(n int) bool {
	var i, j, k int
	i = n % 10
	j = (n / 10) % 10
	k = (n / 10 / 10) % 10

	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}

func isShuiXianHua(str string) bool {
	var result = 0
	for i := 0; i < len(str); i++ {
		num := int(str[i] - '0')
		result += (num * num * num)
	}

	number, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if result == number {
		return true
	} else {
		return false
	}
}

func sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		temp := 1
		for j := 0; j < i; j++ {
			temp = temp * i
		}
		fmt.Printf("%d!=%d\n", i, temp)
		sum += temp
	}
	return sum
}

func main() {
	//var m int
	//
	//fmt.Scanf("%d %d", &n, &m)
	//for i:= n; i <= m; i++ {
	//	if isPrim(i) {
	//		fmt.Printf("%d是素数\n", i)
	//	} else {
	//		fmt.Printf("%d是NO素数\n", i)
	//	}
	//}
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(str)
	//fmt.Println(isShui(n))
	fmt.Println(isShuiXianHua(str))
}
