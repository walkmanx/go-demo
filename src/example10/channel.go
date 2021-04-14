package main

import (
	"fmt"
)

func main() {

	rangeFun()
}

// 通过for循环遍历channel的内容
func forechFun() {

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

// 通过range遍历channel内容
func rangeFun() {
	//创建无缓冲channel
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子线程 i=", i)
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println("num=", num)
	}
}
