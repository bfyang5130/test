package protocol

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)


//处理配置文件的内容转化为一个数组，可以很好地把拿数据拿出来进行处理
func Readconfig() (strErr error,targetPath string){

	//读取配置文件
	config, err := yaml.ReadFile("../config.yaml")
	if err != nil {
		return fmt.Errorf("%s","读取配置文件失败"),""
	}
	//获得源文件路径
	targetPath,err=config.Get("target_path")
	if err!=nil{
		return fmt.Errorf("%s","源文件路径配置不正确"),""
	}
	fmt.Println(fmt.Sprintf("同步目录为:%s",targetPath))
	return nil,targetPath
}