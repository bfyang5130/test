package main

import (
	"net"
	"fmt"
)

func main(){
	//为了支持热配置，所以一开始不读取配置文件，当监控到文件变动时，才读取配置进行传送
	address := `127.0.0.1:60010`
	//开始连接服务器
	fmt.Printf("开始创建TCP连接,连接到：%s\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("连接失败:%s\n", err.Error())
		return
	}
	conn.Write([]byte("adfdfdfdfdf"))
	defer conn.Close()
}