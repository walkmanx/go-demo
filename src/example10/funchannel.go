package main

import (
	"fmt"
	"time"
)

func main() {

	//创建无缓冲channel
	ch := make(chan string)

	defer fmt.Println("主线程结束")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子线程 i=", i)
			time.Sleep(time.Second)
		}
		ch <- "子线程执行结束"
	}()

	// 接收ch数据并且赋值给str，如果ch没有数据则会阻塞
	str := <-ch
	fmt.Println("str = ", str)

}
