package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func writeFile(title, content string) error {
	filename := title + ".html"
	return ioutil.WriteFile(filename, []byte(content), 0600)
}

func HttpGet(url string) (result string, err error) {
	resp, geterr := http.Get(url)
	if geterr != nil {
		err = geterr
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 1024*4)

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err : ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func DoSearch(start, end int) {
	fmt.Printf("开始查询从第 %d 页到第 %d 页的数据\n", start, end)

	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=go&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url - :", url)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet error ：", err)
			continue
		}
		werr := writeFile(strconv.Itoa(i), result)
		if werr != nil {
			fmt.Println("写入文件出错:", werr)
			continue
		}
	}
}

func main() {
	var start, end int

	fmt.Printf("请输入起始页 （ >= 1）:")
	fmt.Scan(&start)

	fmt.Printf("请输入终止页 （ >= 起始页）:")
	fmt.Scan(&end)

	DoSearch(start, end)
}
