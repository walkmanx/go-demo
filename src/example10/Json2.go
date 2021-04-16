package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	//Company不会导出到JSON中
	Company string `json:"-"`

	// Subjects 的值会进行二次JSON编码
	Subjects []string `json:"subjects"`

	//转换为字符串，再输出
	IsOk bool `json:",string"`

	// 如果 Price 为空，则不输出到JSON串中
	Price float64 `json:"price, omitempty"`
}

func TestJson1() {
	t1 := IT{Company: "itcast", Subjects: []string{"Go", "C++", "Python", "Test"}, IsOk: true}

	b, err := json.Marshal(t1)
	//json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//输出结果：{"subjects":["Go","C++","Python","Test"],"IsOk":"true","price":0}
}

func TestMapToJson() {
	// 创建一个保存键值对的映射
	t1 := make(map[string]interface{})
	t1["company"] = "itcast"
	t1["subjects "] = []string{"Go", "C++", "Python", "Test"}
	t1["isok"] = true
	t1["price"] = 666.666

	b, err := json.Marshal(t1)
	//json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//输出结果：{"company":"itcast","isok":true,"price":666.666,"subjects ":["Go","C++","Python","Test"]}
}

func main() {
	TestJson1()

	TestMapToJson()
}
