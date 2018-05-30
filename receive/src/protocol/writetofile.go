package protocol

import (
	"fmt"
	"os"
	"strings"
)

func WriteToFile(buf []byte,targetPath string) {
	/*
	//获得要传输的文件路径
	if fileList.Len()<1{
		fmt.Println("没有要传输的文件路径")
		return
	}
	filePath:=fileList.Back()

	newFilePath:=fmt.Sprintf("%s",filePath.Value)
	fmt.Println(fmt.Sprintf("我要传输到:%s",newFilePath))
	f, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
	if err!=nil{
		fmt.Println("无法打开要传输的文件")
		return
	}
	defer f.Close()
	*/
	isString:=string(buf)
	//判断是否为传输文件
	if strings.HasPrefix(isString, "qystofile///") && strings.Index(isString,"qyspath///")>=0 {
        //把传输路径拿出来
        fmt.Println(strings.Index(isString,"qyspath///"))
        theNewPath:=isString[13:strings.Index(isString,"qyspath///")]
        if !strings.HasPrefix(theNewPath,"/"){
        	theNewPath=fmt.Sprintf("/%s",theNewPath)
		}
        //输出新路径
        fmt.Println(theNewPath)
        //配置目标真实路径
        theNewTargetPath:=fmt.Sprintf("%s%s",targetPath,theNewPath)
        fmt.Println(theNewTargetPath)
        //打开文件
		f, err := os.OpenFile(theNewTargetPath, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
		if err!=nil{
			fmt.Println("无法打开要传输的文件")
			return
		}
		defer f.Close()
        //把写入的信息去掉路径标识
        newBufString:=isString[strings.Index(isString,"qyspath///")+10:]


        f.Write([]byte(newBufString))
	}
}
//前提本方法不做判断处理，进来的必须保证为有效的路径
func fitFilePath(oldPath string,rePath string)(newPath string){
	newPath=strings.Replace(oldPath, "\\", "/", -1)
	newPath=strings.Replace(newPath,rePath,"",1)
	return newPath
}