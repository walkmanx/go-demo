package main

import (
	"fmt"
	"time"
)

// 定义一个全局channel变量
var ch = make(chan int)

// 打印机（将一个字符串按照每个字符打印）
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	// 发送666到ch
	ch <- 666
	fmt.Printf("\n")
}

func Person1() {
	Printer("hello")
}

func Person2() {
	// 接收数据，如果ch没有值则会阻塞
	<-ch
	Printer("world")
}

func main() {
	// 通过channel实现同步和数据交互
	go Person1()
	go Person2()

	for {
	}

}
