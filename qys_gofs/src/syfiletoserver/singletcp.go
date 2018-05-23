package syfiletoserver

import (
	"fmt"
	"net"
)

type singleTcp struct {
	conn net.Conn
}

// private
var instance *singleTcp

// public
func GetInstance(address string) *singleTcp {
	if instance == nil {
		//开始连接服务器
		fmt.Println(fmt.Sprintf("开始创建TCP连接,连接到：%s", address))
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println(fmt.Sprintf("连接失败：%s"), err.Error())
			instance=nil
		}
		defer conn.Close()
		instance = &singleTcp{conn}     // not thread safe
	}
	fmt.Println(fmt.Sprintf("连接到：%s", address))
	return instance
}