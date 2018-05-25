package main

import "github.com/fsnotify/fsnotify"
import (
	"syfiletoserver"
	"fmt"
)

func main() {
	//获取配置文件
	err,sourcePath,reCon:=syfiletoserver.Readconfig()
	sourcePath=fmt.Sprintf("%s",sourcePath)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//创建发送记操作日志
	err=syfiletoserver.CreateLogFile("main.bin")
	if err!=nil{
		fmt.Println("主同步日志无法创建，程序无法运行")
		return
	}
	//开始监控文件
	//-------------监控开始-----------------

	watch, _ := fsnotify.NewWatcher()
	w := syfiletoserver.Watch{
		watch,
	}
	fmt.Println(sourcePath)
	w.WatchDir(sourcePath,reCon);
	select {};
}