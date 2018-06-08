package protocol

import (
	"os"
	"fmt"
	"container/list"
	"path"
)

//创建文件
func CreateFile(op bool,targetPath string,opPath string,fileList *list.List) (error) {
	//默认记录日志就在log目录下面
	//这里后期要对是目录的进行处理，目前就是如果存在先删除，再创建
	newPath := fmt.Sprintf("%s%s", targetPath, opPath)
	//死活不管，先把目录创建出来
	/**
	dirPostion:=strings.LastIndex(newPath,`/`)
	if !(dirPostion<0){
		newPath=newPath[:dirPostion]
	}
	*/
	if err:=CreateDir(newPath,``,fileList); err!=nil{
		return fmt.Errorf("创建文件目录失败")
	}
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

//创建目录
func CreateDir(targetPath string,opPath string,fileList *list.List)(error){
	//上层直接处理掉不做判断
	newPath := fmt.Sprintf("%s%s", targetPath, opPath)
	fmt.Println("准备创建",path.Dir(newPath))
	err := os.MkdirAll(path.Dir(newPath),0666) //创建目录
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("创建目录失败")
	}
	return nil
	//默认记录日志就在log目录下面
	//这里后期要对是目录的进行处理，目前就是如果存在先删除，再创建
	if !checkIsDir(newPath) {
		firstE:=fileList.Front()
		if firstE!=nil && firstE.Value==newPath {
			return nil
		}
		fileList.PushFront(newPath)
		firstE=fileList.Front()
		err := os.MkdirAll(path.Dir(newPath),0666) //创建目录
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("创建目录失败")
		}
		fmt.Println("成功创建目录",newPath)
		fileList.Remove(firstE)
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
/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkIsDir(filename string) bool {
	var exist = false
	fileInfo, err := os.Stat(filename)
	if (os.IsNotExist(err)) {
		fmt.Println("不存这个目录",filename)
		return false
	}
	if (fileInfo.IsDir()) {
		//传输一个创建目录的指令过去
		fmt.Println("这个目录已经创建",filename)
		return true
	}
	return exist
}