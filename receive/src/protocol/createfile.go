package protocol

import (
	"os"
	"fmt"
)

func CreateFile(op bool,targetPath string,opPath string) (error) {
	//默认记录日志就在log目录下面
	//这里后期要对是目录的进行处理，目前就是如果存在先删除，再创建
	newPath := fmt.Sprintf("%s%s", targetPath, opPath)
	if !checkFileIsExist(newPath) {
		_, err := os.Create(newPath) //创建文件
		if err != nil {
			return fmt.Errorf("创建文件失败")
		}
	}else{
		err:=os.Remove(newPath)
		if err !=nil{
			return fmt.Errorf("删除文件失败")
		}
		_, err = os.Create(newPath) //创建文件
		if err != nil {
			return fmt.Errorf("创建文件失败")
		}
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