package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer conn.Close()

	// 接收服务器回复的数据
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err = ", err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	// 从键盘输入内容向服务器发送消息
	str := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("os.Stdin.Read err = ", err)
			return
		}
		conn.Write(str[:n])
	}
}
