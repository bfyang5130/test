package syfiletoserver

import (
	"fmt"
	"net"
)

func SyFileToServer(opType string,filePath string,serverPort string)(error){
	//获得连接
	//开始连接服务器
	fmt.Println(fmt.Sprintf("开始创建TCP连接,连接到：%s", serverPort))
	conn, err := net.Dial("tcp", serverPort)
	if err != nil {
		return fmt.Errorf("连接失败")
	}
	defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("%s/%s",opType,filePath)))
	   return fmt.Errorf("传输失败")
}