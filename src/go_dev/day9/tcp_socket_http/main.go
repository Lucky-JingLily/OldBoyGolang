package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\n"
	msg += "Host:www.baidu.com\n"
	msg += "Connection:keep-alive\n"
	msg += "\n"

	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed,", err)
		return
	}
	fmt.Println("send to baidu.com bytes", n)
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		fmt.Println("count:", count, " err:", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}
