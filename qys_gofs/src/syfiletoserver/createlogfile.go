package syfiletoserver

import (
	"os"
	"fmt"
)

func CreateLogFile(fileName string) (error){
	//默认记录日志就在log目录下面
	newPath:=fmt.Sprintf("../log/%s",fileName)
	if !checkFileIsExist(newPath) {
		_, err := os.Create(newPath) //创建文件
		if err!=nil {
			return fmt.Errorf("创建文件失败")
		}
		return nil
	}
    return nil
}
/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*
 *重新检测传输过去的是否正确
 */
func ReCheckOpFile(baseFilePath string,opType string,newPath string) (optype string,newpath string) {
	fileinfo, err := os.Stat(baseFilePath)
	if err==nil{
		if fileinfo.IsDir(){
			return "C",fmt.Sprintf("%s/",newPath)
		}
	}
	return opType,newPath
}