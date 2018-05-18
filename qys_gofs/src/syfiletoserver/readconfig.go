package syfiletoserver

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)

func Readconfig(){

	//读取配置文件
	config, err := yaml.ReadFile("../config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	//获得源文件路径
	sourcePath,err:=config.Get("source_path")
	if err!=nil{
		fmt.Println("源文件路径配置不正确")
	}
	fmt.Println(fmt.Printf("同步源文件为:%s",sourcePath))
	//获得目标路径
	targetPath,err:=config.Get("target_path")
	if err!=nil{
		fmt.Println("目标路径配置不正确")
	}
	fmt.Println(fmt.Printf("同步源文件为:%s",targetPath))
	//获得同步服务器的IP与port
	newNode,err:=yaml.Child(config.Root,"servers")
	if err != nil{
		fmt.Println(err)
	}
	lst,ok := newNode.(yaml.List)
	if !ok {
		fmt.Println("列表配置有误")
	}
	for _,v := range lst{
		v1,ok:=v.(yaml.Map)
		if !ok {
			fmt.Printf("配置结构有错")
		}
		fmt.Printf("IP:%s,PORT:%s",v1["server"],v1["port"])
	}
}