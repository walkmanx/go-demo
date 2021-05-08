package main

import (
	"fmt"

	"example14/model"
)

func main() {
	model.setName("张三")
	fmt.Println(model.getName())
}
