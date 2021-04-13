package main

import (
	"fmt"

	"com.zy/greeting"
)

func main() {
	message := greeting.Greet("张三")
	fmt.Println(message)
}
