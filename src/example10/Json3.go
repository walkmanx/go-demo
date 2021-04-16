package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	b := []byte(`{
	    "company": "itcast",
	    "subjects": [
	        "Go",
	        "C++",
	        "Python",
	        "Test"
	    ],
	    "isok": true,
	    "price": 666.666
	}`)

	var t IT
	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(t)
	//运行结果：{itcast [Go C++ Python Test] true 666.666}

	//只想要Subjects字段
	type IT2 struct {
		Subjects []string `json:"subjects"`
	}

	var t2 IT2
	err = json.Unmarshal(b, &t2)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(t2)
	//运行结果：{[Go C++ Python Test]}
}
