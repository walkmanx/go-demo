package main

import (
	"fmt"
	"time"
)

// 打印机（将一个字符串按照每个字符打印）
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("hello")
}

func Person2() {
	Printer("world")
}

func main() {

	go Person1()
	go Person2()

	for {

	}
}
