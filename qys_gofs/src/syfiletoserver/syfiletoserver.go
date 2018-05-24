package syfiletoserver

import (
	"fmt"
	"net"
	"protocol"
	"bufio"
	"io"
	"os"
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
	   //进行封包操作传递命令
	   conn.Write(protocol.Packet([]byte(fmt.Sprintf("%s/%s",opType,filePath))))
	   //对新建文件写入文件，修改文件进行文件传输处理
	   switch opType{
	   case "c":
		   //readBufio(filePath,conn)
	       break;
	   case "w":
		   readBufio(filePath,conn)
	       break;
	   default:
		   fmt.Println("iamhere3")
		   break;
	   }
	   return fmt.Errorf("传输失败")
}
//读取并写入
func readBufio(path string,conn net.Conn) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开目录")
		return
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("无法读取文件")
		}
		if 0 == readNum {
			break
		}
		conn.Write(protocol.Packet(buf[:readNum]))
	}
}