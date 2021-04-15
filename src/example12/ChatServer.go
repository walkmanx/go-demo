package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string      // 用户名
	Addr string      // IP地址
	C    chan string // 通信通道
}

// 在线用户列表
var onLineUsers map[string]User

var message = make(chan string)

// 给在线用户广播消息
func Broadcast() {

	onLineUsers = make(map[string]User)

	for {
		msg := <-message

		for _, v := range onLineUsers {
			v.C <- msg
		}

	}

}

// 给客户端发送消息
func SendMessage(user User, conn net.Conn) {
	for msg := range user.C {
		conn.Write([]byte(msg + "\n"))
	}
}

// 组装固定格式的消息
func BuildMessage(u User, msg string) (buf string) {
	buf = "[" + u.Addr + "]" + u.Name + ":" + msg + "\n"
	return
}

// 处理用户请求保存用户信息在onLineUsers
func HandlerConn(conn net.Conn) {

	defer conn.Close()

	// 获取客户端IP地址
	addr := conn.RemoteAddr().String()

	u := User{addr, addr, make(chan string)}

	onLineUsers[addr] = u

	go SendMessage(u, conn)

	// 给所有在线用户发送上线消息
	message <- BuildMessage(u, "上线")

	// 接收客户端发送的数据
	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)

			if n == 0 {
				fmt.Println("conn.Read err :", err)
				return
			}

			msg := string(buf[:n-2])

			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list : \n"))

				for _, v := range onLineUsers {
					conn.Write([]byte(BuildMessage(v, "")))
				}

			} else {
				message <- BuildMessage(u, msg)
			}
		}
	}()

	for {

	}

}

func main() {

	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}

	defer listener.Close()

	go Broadcast()

	for {
		conn, connerr := listener.Accept()
		if connerr != nil {
			fmt.Println("connerr = ", connerr)
			return
		}

		go HandlerConn(conn)

	}

}
