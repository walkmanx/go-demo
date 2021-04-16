package main

import (
	"fmt"
)

func swap01(a, b int) {
	a, b = b, a
	fmt.Printf("swap01 a = %d, b = %d\n", a, b)
}

func swap02(x, y *int) {
	*x, *y = *y, *x
}

// 指针做函数参数
func main() {
	a := 10
	b := 20

	swap01(a, b) //值传递
	//swap02(&a, &b) //变量地址传递
	fmt.Printf("a = %d, b = %d\n", a, b)
}
