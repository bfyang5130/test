package syfiletoserver

import (
	"fmt"
	"net"
	"protocol"
	"os"
	"io/ioutil"
)

func SyFileToServer(opType string, filePath string, newPath string, serverPort string) (error) {
	//在这里转过来的filePath混乱出错，所以在这里做一次目录检查
	opType,newPath=ReCheckOpFile(filePath,opType,newPath)
	//获得连接
	//开始连接服务器
	fmt.Println(fmt.Sprintf("开始创建TCP连接,连接到：%s", serverPort))
	conn, err := net.Dial("tcp", serverPort)
	if err != nil {
		return fmt.Errorf("连接失败")
	}
	defer conn.Close()
	//进行封包操作传递命令
	conn.Write(protocol.Packet([]byte(fmt.Sprintf("qyssyfile///%s/%s", opType, newPath))))
	//对新建文件写入文件，修改文件进行文件传输处理
	switch opType {
	case "c":
		//readBufio(filePath,conn)
		break;
	case "w":
		fmt.Println("i am wrinting here")
		readBufio(filePath, newPath, conn)
		break;
	default:
		fmt.Println("iamhere3")
		break;
	}
	return fmt.Errorf("传输失败")
}

//读取并写入
//因为读取数据过程中，会一次读入内存，所以限制了2m大的文件，但是一次多个文件读取时有没有问题还要做处理
func readBufio(path string, newPath string, conn net.Conn) {
	//读取文件信息，如果文件大于2m的话，那么直接不做传输
	fileInfo, err := os.Stat(path)
	if (os.IsNotExist(err)) {
		fmt.Println("文件无法打开")
		return
	}
	if (fileInfo.IsDir()) {
		//传输一个创建目录的指令过去
		conn.Write(protocol.Packet([]byte(fmt.Sprintf("qyssyfile///C/%s", newPath))))
		fmt.Println("目录不需传输")
		return
	}
	filesize := fileInfo.Size()
	//暂时不处理大文件传输
	if (filesize > 2*1024*1024) {
		fmt.Println("文件大于2m不做传输")
		return
	}
	//-------------------------start openfile all------------
	fileobj, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if (err != nil) {
		fmt.Println("无法打开文件")
		return
	}
	buf, err := ioutil.ReadAll(fileobj)
	if (err != nil) {
		fmt.Println("读取文件失败")
	}

	defer fileobj.Close()
	//-------------------------end openfile all--------------
	/////////////////////////////////////////////////////////////

	//处理新路径配置
	splitPath := fmt.Sprintf("qystofile///%sqyspath///", newPath)
	conn.Write(protocol.Packet(append([]byte(splitPath), buf...)))

}
