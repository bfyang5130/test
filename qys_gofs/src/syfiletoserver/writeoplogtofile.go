package syfiletoserver

import (
	"fmt"
	"os"
	"io"
	"github.com/satori/go.uuid"
)

func WriteOpLogFile(opType string,filePath string,reCon []SerConfig){
	//获得配置信息
	//重新读取一次配置文件信息，保证能够热更新
	err,reCon2:=Readconfig()
	if err!=nil{
		fmt.Println("新修改的文件不正确，无法热更新配置文件，请重新配置")
	}
	//写操作日志到主文件
     writeLogTofile(opType,filePath,"main.bin")
	//如果配置没有问题，重新赋值。
	reCon=reCon2
	//循环处理配置信息
	for _,v:=range reCon {
		//读取同步文件
		//拼接文件
		newFileName:=fmt.Sprintf("../log/%s%s",v.server,v.port)
		err=CreateLogFile(newFileName)
		if err==nil{
			//同步文件到各个端
			err:=SyFileToServer(opType,filePath,fmt.Sprintf("%s:%s",v.server,v.port))
			//写入操作命令到各自的同步文件
			if err==nil{
			writeLogTofile(opType,filePath,fmt.Sprintf("%s%s",v.server,v.port))
			}
		}
	}
}

func writeLogTofile(opType string,opFilePath string,logFile string) {
	f, err := os.OpenFile(fmt.Sprintf("../log/%s",logFile), os.O_APPEND, 0666) //打开文件
	if err!=nil{
		fmt.Println("无法打开同步日志文件")
		return
	}
	defer f.Close()
	//生成一个唯一串
	u1 := uuid.Must(uuid.NewV4())
	_, err1 := io.WriteString(f, fmt.Sprintf("%s\n", u1)) //写入文件(字符串)
	if err1==nil{
		_, err1 = io.WriteString(f, fmt.Sprintf("%s %s\n", opType,opFilePath))
	}
}