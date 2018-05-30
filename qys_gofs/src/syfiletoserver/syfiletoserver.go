package syfiletoserver

import (
	"fmt"
	"net"
	"protocol"
	"bufio"
	"io"
	"os"
)

func SyFileToServer(opType string,filePath string,newPath string,serverPort string)(error){
	//获得连接
	//开始连接服务器
	fmt.Println(fmt.Sprintf("开始创建TCP连接,连接到：%s", serverPort))
	conn, err := net.Dial("tcp", serverPort)
	if err != nil {
		return fmt.Errorf("连接失败")
	}
	defer conn.Close()
	   //进行封包操作传递命令
	   conn.Write(protocol.Packet([]byte(fmt.Sprintf("qyssyfile///%s/%s",opType,newPath))))
	   //对新建文件写入文件，修改文件进行文件传输处理
	   switch opType{
	   case "c":
		   //readBufio(filePath,conn)
	       break;
	   case "w":
		   readBufio(filePath,newPath,conn)
	       break;
	   default:
		   fmt.Println("iamhere3")
		   break;
	   }
	   return fmt.Errorf("传输失败")
}
//读取并写入
func readBufio(path string,newPath string,conn net.Conn) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开目录")
		return
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	////////////////////////////////////////////////////////////

	fl, err := os.OpenFile("E:/index.html", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fl.Close()

	/////////////////////////////////////////////////////////////

	//处理新路径配置
	splitPath:=fmt.Sprintf("qystofile///%sqyspath///",newPath)
	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("无法读取文件")
		}
		if 0 == readNum {
			//读取文件结尾标志
			break
		}
		//这里增加一个文件目的，才能知道这传输的内容是到那个文件
		n, err1 := fl.Write(append([]byte(splitPath),buf[:readNum]...))
		if err1 == nil && n < len(append([]byte(splitPath),buf[:readNum]...)) {

		}
		conn.Write(protocol.Packet(append([]byte(splitPath),buf[:readNum]...)))

	}
}