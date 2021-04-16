package main

import (
	"fmt"
)

func main() {
	// 切片的创建和初始化
	var s1 []int //声明切片和声明array一样，只是少了长度，此为空(nil)切片
	s2 := []int{}

	//make([]T, length, capacity) //capacity省略，则和length的值相同
	var s3 []int = make([]int, 0)
	s4 := make([]int, 0, 0)

	s5 := []int{1, 2, 3} //创建切片并初始化

	// 切片和底层数组关系
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[2:5]       //[2 3 4]
	s1[2] = 100        //修改切片某个元素改变底层数组
	fmt.Println(s1, s) //[2 3 100] [0 1 2 3 100 5 6 7 8 9]

	s2 := s1[2:6] // 新切片依旧指向原底层数组 [100 5 6 7]
	s2[3] = 200
	fmt.Println(s2) //[100 5 6 200]

	fmt.Println(s) //[0 1 2 3 100 5 6 200 8 9]

	// append函数向 slice 尾部添加数据，返回新的 slice 对象
	var s1 []int //创建nil切换
	//s1 := make([]int, 0)
	s1 = append(s1, 1)       //追加1个元素
	s1 = append(s1, 2, 3)    //追加2个元素
	s1 = append(s1, 4, 5, 6) //追加3个元素
	fmt.Println(s1)          //[1 2 3 4 5 6]

	s2 := make([]int, 5)
	s2 = append(s2, 6)
	fmt.Println(s2) //[0 0 0 0 0 6]

	s3 := []int{1, 2, 3}
	s3 = append(s3, 4, 5)
	fmt.Println(s3) //[1 2 3 4 5]

	// 函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准，两个 slice 可指向同一底层数组
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:] //{8, 9}
	s2 := data[:5] //{0, 1, 2, 3, 4}
	copy(s2, s1)   // dst:s2, src:s1

	fmt.Println(s2)   //[8 9 2 3 4]
	fmt.Println(data) //[8 9 2 3 4 5 6 7 8 9]

}
