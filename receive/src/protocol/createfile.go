package protocol

import (
	"os"
	"fmt"
	"container/list"
)

func CreateFile(op bool,targetPath string,opPath string,fileList *list.List) (error) {
	//默认记录日志就在log目录下面
	newPath := fmt.Sprintf("%s%s", targetPath, opPath)
	if !checkFileIsExist(newPath) {
		_, err := os.Create(newPath) //创建文件
		if err != nil {
			return fmt.Errorf("创建文件失败")
		}
	}
	if op {
	if fileList.Len() > 0 {
		fileList.Remove(fileList.Back())
	}
		fileList.PushFront(newPath)
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