package main

import (
	"fmt"
	"time"
)

func main() {
	// Ticker是一个定时触发的计时器，它会以一个间隔(interval)往channel发送一个事件(当前时间)，
	// 而channel的接收者可以以固定的时间间隔从channel中读取事件
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println("i = ", i)

		if i == 5 {
			ticker.Stop()
			break
		}
	}
}
