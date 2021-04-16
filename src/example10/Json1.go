package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func TestJson1() {
	t1 := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//生成一段JSON格式的文本
	//如果编码成功， err 将赋于零值 nil，变量b 将会是一个进行JSON格式化之后的[]byte类型
	b, berr := json.Marshal(t1)
	//输出结果：{"Company":"itcast","Subjects":["Go","C++","Python","Test"],"IsOk":true,"Price":666.666}

	if berr != nil {
		fmt.Println("json berr:", berr)
	}
	fmt.Println(string(b))
	fmt.Println("\n")

	c, cerr := json.MarshalIndent(t1, "", "    ")
	/*
	   输出结果：
	   {
	       "Company": "itcast",
	       "Subjects": [
	           "Go",
	           "C++",
	           "Python",
	           "Test"
	       ],
	       "IsOk": true,
	       "Price": 666.666
	   }
	*/
	if cerr != nil {
		fmt.Println("json cerr:", cerr)
	}
	fmt.Println(string(c))
}

func main() {
	TestJson1()
}
