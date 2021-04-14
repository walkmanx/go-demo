package main

import (
	"fmt"
)

func main() {

	//创建无缓冲channel
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子线程 i=", i)
			ch <- i
		}
		close(ch)
	}()

	for {
		// 如果ok为true则管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else {
			fmt.Println("管道关闭")
			break
		}
	}
}
