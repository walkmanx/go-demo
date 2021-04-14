package main

import (
	"fmt"
)

// 生成者，只能写
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

//消费者，只能读
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}
