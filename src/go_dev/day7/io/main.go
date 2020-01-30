package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := "hello world"
	file, err := os.OpenFile("tmp.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error:", err)
	}
	fmt.Fprintf(file, str)
	file.Close()

	reader := bufio.NewReader(os.Stdin)
	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string fail")
		return
	}
	fmt.Println(str)
}
