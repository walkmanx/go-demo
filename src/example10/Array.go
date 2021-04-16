package main

func main() {
	var n int = 10
	// 数组长度必须是常量，且是类型的组成部分
	var a [n]int  //err, non-constant array bound n
	var b [10]int //ok

	// 操作数组
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 1
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	//range具有两个返回值，第一个返回值是元素的数组下标，第二个返回值是元素的值
	for i, v := range a {
		fmt.Println("a[", i, "]=", v)
	}

	// 内置函数 len(长度) 和 cap(容量) 都返回数组⻓度
	a := [10]int{}
	fmt.Println(len(a), cap(a)) //10 10

	a := [3]int{1, 2}           // 未初始化元素值为 0
	b := [...]int{1, 2, 3}      // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200} // 通过索引号初始化元素，未初始化元素值为 0
	fmt.Println(a, b, c)        //[1 2 0] [1 2 3] [0 0 100 0 200]

	//支持多维数组
	d := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	e := [...][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}} //第二维不能写"..."
	f := [4][2]int{1: {20, 21}, 3: {40, 41}}
	g := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(d, e, f, g)

	// 相同类型的数组之间可以使用 == 或 != 进行比较，但不可以使用 < 或 >，也可以相互赋值
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2}
	fmt.Println(a == b, b == c) //true false

	var d [3]int
	d = a
	fmt.Println(d) //[1 2 3]

}
