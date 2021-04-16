package main

import (
	"fmt"
	"runtime"
)

// 用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行
func Gosched() {
	//创建一个goroutine
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched() //import "runtime"
		/*
		   屏蔽runtime.Gosched()运行结果如下：
		       hello
		       hello

		   没有runtime.Gosched()运行结果如下：
		       world
		       world
		       hello
		       hello
		*/
		fmt.Println("hello")
	}
}

// 调用 runtime.Goexit() 将立即终止当前 goroutine 执行，调度器确保所有已注册 defer延迟调用被执行
func Goexit() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 终止当前 goroutine, import "runtime"

			fmt.Println("B") // 不会执行
		}()

		fmt.Println("A") // 不会执行
	}() //别忘了()

	//死循环，目的不让主goroutine结束
	for {
	}
}

func main() {
	Gosched()
}
