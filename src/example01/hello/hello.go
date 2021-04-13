package main

import (
	"fmt"

	"com.zy/greeting"
)

func main() {
	message := greeting.Hello("张三")
	fmt.Println(message)
}
