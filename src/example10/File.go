package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile() {
	fout, err := os.Create("./xxx.txt") //新建文件
	//fout, err := os.OpenFile("./xxx.txt", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fout.Close() //main函数结束前， 关闭文件

	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\n", "Hello go", i)
		fout.WriteString(outstr)     //写入string信息到文件
		fout.Write([]byte("abcd\n")) //写入byte类型的信息到文件
	}
}

func ReadFile() {
	fin, err := os.Open("./xxx.txt") //打开文件
	if err != nil {
		fmt.Println(err)
	}
	defer fin.Close()

	buf := make([]byte, 1024) //开辟1024个字节的slice作为缓冲
	for {
		n, _ := fin.Read(buf) //读文件
		if n == 0 {           //0表示已经到文件结束
			break
		}

		fmt.Println(string(buf)) //输出读取的内容
	}
}

func CopyFile() {
	args := os.Args //获取用户输入的所有参数

	//如果用户没有输入,或参数个数不够,则调用该函数提示用户
	if args == nil || len(args) != 3 {
		fmt.Println("useage : xxx srcFile dstFile")
		return
	}

	srcPath := args[1] //获取输入的第一个参数
	dstPath := args[2] //获取输入的第二个参数
	fmt.Printf("srcPath = %s, dstPath = %s\n", srcPath, dstPath)

	if srcPath == dstPath {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	srcFile, err1 := os.Open(srcPath) //打开源文件
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	dstFile, err2 := os.Create(dstPath) //创建目的文件
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	buf := make([]byte, 1024) //切片缓冲区
	for {
		//从源文件读取内容，n为读取文件内容的长度
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		if n == 0 {
			fmt.Println("文件处理完毕")
			break
		}

		//切片截取
		tmp := buf[:n]
		//把读取的内容写入到目的文件
		dstFile.Write(tmp)
	}

	//关闭文件
	srcFile.Close()
	dstFile.Close()
}

func main() {
	// WriteFile()

	// ReadFile()

	CopyFile()
}
