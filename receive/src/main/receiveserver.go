package main

import (
	"fmt"
	"net"
	"os"
	"../protocol"
	"strings"
	"container/list"
)

func main() {
	//获取配置文件
	err,targetPath:=protocol.Readconfig()
	targetPath=fmt.Sprintf("%s",targetPath)
	if err!=nil{
		fmt.Println("没有配置同步目录")
		return
	}
	//定义一个list来装来处理的文件
	fileList:=list.New()
	netListen, err := net.Listen("tcp", "127.0.0.1:60010")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn,targetPath,fileList)
	}
}

func handleConnection(conn net.Conn,targetPath string,fileList *list.List) {

	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)
	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel,targetPath,fileList)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		/*        Log(conn.RemoteAddr().String(), "receive data length:", n)
				Log(conn.RemoteAddr().String(), "receive data:", buffer[:n])
				Log(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))
		*/
		tmpBuffer = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}
func reader(readerChannel chan []byte,targetPath string,fileList *list.List) {
	for {
		select {
		case data := <-readerChannel:
			//判断是否为命令
			isString:=string(data)
			if strings.HasPrefix(isString, "qyssyfile///"){

				isReplace:=strings.Replace(isString,"qyssyfile///","",1)
				//载掉第一个字符串
				opType:=isReplace[:1]
				fmt.Println(fmt.Sprintf("这是一个操作命令:%s",opType))
				opPath:=isReplace[2:]
				switch opType {
				case "c":
					//创建文件
					protocol.CreateFile(true,targetPath,opPath)
					break
				case "w":
					//写入文件
					protocol.CreateFile(false,targetPath,opPath)
					break
				default:
					break
				}
			}else {
				//写入文件
				protocol.WriteToFile(data,targetPath)
			}

		}
	}
}
func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}