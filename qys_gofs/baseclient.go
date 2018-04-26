package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	address := `127.0.0.1:60010`
	//开始连接服务器
	fmt.Printf("开始创建TCP连接,连接到：%s\n", address)
	conn, err := net.Dial("tcp", address)
	defer conn.Close()
	if err != nil {
		fmt.Printf("连接失败:%s\n", err.Error())
	}
	fmt.Println("连接成功")
	fmt.Println("开始传送数据...")

	newFile:=Read1()
	//读取文件里的东西
	conn.Write(newFile)
}

func Read1() []byte{
	path:="F:/test/flip.html"
	fi,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte,1024,1024)
	buf := make([]byte,1024)
	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return chunks
}