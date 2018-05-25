package protocol

import (
	"fmt"
	"os"
	"strings"
	"container/list"
)

func WriteToFile(buf []byte,fileList *list.List) {
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
	//写入内容
	f.Write(buf)
}
//前提本方法不做判断处理，进来的必须保证为有效的路径
func fitFilePath(oldPath string,rePath string)(newPath string){
	newPath=strings.Replace(oldPath, "\\", "/", -1)
	newPath=strings.Replace(newPath,rePath,"",1)
	return newPath
}