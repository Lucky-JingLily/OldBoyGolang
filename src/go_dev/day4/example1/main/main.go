package main

import (
	"fmt"
	"strings"
)

func count(str string) (wordCount, spaceCount, numCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numCount++
		default:
			otherCount++
		}
	}
	return
}

func add(str1, str2 string) (result string) {
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}

	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left = 0

	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)

		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}

	return
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//result, _, error := reader.ReadLine()
	//if error != nil {
	//	fmt.Println("输入读取失败", error)
	//	return
	//}
	//w, s, n, o :=  count(string(result))
	//fmt.Println(w, s, n, o)

	var str string
	fmt.Scanf("%s", &str)
	temp := strings.Split(str, "+")
	str1 := strings.TrimSpace(temp[0])
	str2 := strings.TrimSpace(temp[1])
	fmt.Println(add(str1, str2))
}
