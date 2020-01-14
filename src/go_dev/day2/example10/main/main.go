package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func reverse(str string) string {
	//var result string
	//strLen := len(str)
	//for i := 0; i < strLen; i++ {
	//	result = result + fmt.Sprintf("%c", str[strLen-i-1])
	//}
	//
	//return result
	//tmp := []byte(str)
	//var result string
	//for i := len(str)-1; i >= 0; i-- {
	//	result = result + fmt.Sprintf("%c", tmp[i])
	//}
	//return result
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}

	return string(result)
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(rand.Int())
		fmt.Println(rand.Intn(100))
		fmt.Println(rand.Float32())
	}

	str := "hello world"
	fmt.Println(reverse(str))
}
