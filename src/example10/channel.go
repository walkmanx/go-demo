package main

import (
	"fmt"
)

func definedChan() {
	// make(chan Type) //等价于make(chan Type, 0)
	// make(chan Type, capacity)

	//当 capacity= 0 时，channel 是无缓冲阻塞读写的，
	//当capacity> 0 时，channel 有缓冲、是非阻塞的，直到写满 capacity个元素才阻塞写入。

	//channel通过操作符<-来接收和发送数据，发送和接收数据语法：
	//channel <- value   //发送value到channel
	//<-channel          //接收并将其丢弃
	//x := <-channel     //从channel中接收数据，并赋值给x
	//x, ok := <-channel //功能同上，同时检查通道是否已关闭或者是否为空
}

func BlockChan() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子协程结束")

		fmt.Println("子协程正在运行……")

		c <- 666 //666发送到c
	}()

	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main协程结束")
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

func main() {

	BlockChan()
}
