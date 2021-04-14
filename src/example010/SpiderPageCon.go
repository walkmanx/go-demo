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

func SpiderPage(url string, i int, finished chan int) {
	fmt.Printf("正在爬取第 %d页数据:", i, url)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet error ：", err)
		return
	}
	werr := writeFile(strconv.Itoa(i), result)
	if werr != nil {
		fmt.Println("写入文件出错:", werr)
		return
	}
	// 发送i到finished
	finished <- i
}

func DoSpiderCon(start, end int) {
	fmt.Printf("开始查询从第 %d 页到第 %d 页的数据\n", start, end)
	// 用make函数定义一个int类型的channel
	finished := make(chan int)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=go&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		// 在函数前加go关键字，可创建并发执行单元
		go SpiderPage(url, i, finished)
	}
	for i := start; i <= end; i++ {
		// channel通过操作符<-来接收和发送数据
		// <-finished 接收并将其丢弃，如果通道没有数据就会阻塞
		fmt.Printf("第%d个页面数据爬起完成\n", <-finished)
	}
}

func main() {
	var start, end int

	fmt.Printf("请输入起始页 （ >= 1）:")
	fmt.Scan(&start)

	fmt.Printf("请输入终止页 （ >= 起始页）:")
	fmt.Scan(&end)

	DoSpiderCon(start, end)
}
