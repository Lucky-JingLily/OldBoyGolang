package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("tmp.log")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line []byte
	for {
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("read file finished")
			break
		}
		if err != nil {
			fmt.Println("readline failed, err:", err)
		}
		if prefix == true {
			// line追加，data...展开追加
			line = append(line, data...)
		} else {
			fmt.Printf("data:%s\n", string(line))
			// 清空line里面数据，不影响读下一行数据
			line = line[:]
		}
	}
}
