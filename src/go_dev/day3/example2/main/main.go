package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func urlProcess(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}

func pathProcess(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	} else {
		return path + "/"
	}
}

func main() {
	//var (
	//	url string
	//	path string
	//)
	//
	//fmt.Scanf("%s%s", &url, &path)
	//url = urlProcess(url)
	//path = pathProcess(path)
	//
	//fmt.Println(url, path)
	start := time.Now().UnixNano()
	str := " hello world abc "
	fmt.Println(strings.Replace(str, "hello", "word", -1))
	fmt.Println(strings.Count(str, "a"))
	fmt.Println(strings.Repeat(str, 3))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Trim(str, " "))
	fmt.Println(strings.TrimLeft(str, " "))
	result := strings.Fields(str)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	result = strings.Split(str, "l")
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	fmt.Println(strings.Join(result, "l"))

	fmt.Println(strconv.Itoa(10000))
	number, error := strconv.Atoi("2000")
	if error != nil {
		fmt.Println("error")
	} else {
		fmt.Println(number)
	}

	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000)
}
