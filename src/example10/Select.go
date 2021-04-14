package main

import (
	"fmt"
)

// 生成斐波那契数列
func main() {
	ch := make(chan int)    // 数字通信
	quit := make(chan bool) // 程序是否结束

	// 消费者，从channel读取内容

	// 生产者，生成数字，写入channel

	go func() {
		for i := 0; i < 20; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()

	fibonacci(ch, quit)

}

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		// 监听channel数据流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("退出计算:", flag)
			return
		}
	}
}
