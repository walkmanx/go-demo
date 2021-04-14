package main

import (
	"fmt"
	"time"
)

func main() {
	reset()
}

func fun1() {
	// 创建一个定时器，设置时间为2秒，往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间:", time.Now())

	// 2秒后，读取timer通道的内容
	t := <-timer.C

	fmt.Println("t:", t)
}

// 延迟2秒执行
func fun2() {

	time.Sleep(2 * time.Second)

	fmt.Println("时间到")

}

// 延迟2秒执行
func fun3() {

	<-time.After(2 * time.Second)

	fmt.Println("时间到")

}

// 停止定时器
func stop() {

	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("定时器时间到")
	}()

	timer.Stop()

	for {

	}
}

// 重置定时器
func reset() {

	timer := time.NewTimer(3 * time.Second)

	ok := timer.Reset(1 * time.Second)

	fmt.Println("ok:", ok)

	<-timer.C

	fmt.Println("时间到")

}
