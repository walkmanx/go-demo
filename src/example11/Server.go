package main

import (
	"fmt"
	"net"
	"strings"
)

// 处理客户端请求
func HandleConn(conn net.Conn) {

	//defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Println("addr connect sucessful :", addr)

	buf := make([]byte, 2048)

	for {

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		fmt.Printf("接收到客户端消息-[%s]：%s\n", addr, string(buf[:n]))

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

	}

}

func main() {

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	for {
		conn, connerr := listener.Accept()
		if connerr != nil {
			fmt.Println("connerr = ", connerr)
			return
		}

		go HandleConn(conn)
	}

}
