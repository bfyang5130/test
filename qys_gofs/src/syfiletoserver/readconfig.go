package syfiletoserver

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)



type SerConfig struct {
	server string;
	port string;
}
//处理配置文件的内容转化为一个数组，可以很好地把拿数据拿出来进行处理
func Readconfig() (strErr error,sourcePath string,reCon []SerConfig){

	//读取配置文件
	config, err := yaml.ReadFile("../config.yaml")
	if err != nil {
		return fmt.Errorf("%s","读取配置文件失败"),"",reCon
	}
	//获得源文件路径
	sourcePath,err=config.Get("source_path")
	if err!=nil{
		return fmt.Errorf("%s","源文件路径配置不正确"),"",reCon
	}
	fmt.Println(fmt.Sprintf("同步源文件为:%s",sourcePath))
	//获得同步服务器的IP与port
	newNode,err:=yaml.Child(config.Root,"servers")
	if err != nil{
		return fmt.Errorf("%s","文件内容格式不正确"),"",reCon
	}
	lst,ok := newNode.(yaml.List)
	if !ok {
		return fmt.Errorf("%s","列表配置有误"),"",reCon
	}
	for _,v := range lst{
		v1,ok:=v.(yaml.Map)
		if !ok {
			return fmt.Errorf("%s","配置结构有错"),"",reCon
		}

		newServer:=fmt.Sprintf("%s",v1["server"])
		newPort:=fmt.Sprintf("%s",v1["port"])
		newSerConfig:=SerConfig{newServer,newPort}
		reCon=append(reCon,newSerConfig)
		//生成记录配置文件
		newFileName:=fmt.Sprintf("%s%s",newServer,newPort)
		err:=CreateLogFile(newFileName)
		if err!=nil{
			fmt.Println(fmt.Sprintf("从机：%s无法创建，无法同步文件"))
		}
	}
	return nil,sourcePath,reCon
}