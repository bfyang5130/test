package main

import "github.com/fsnotify/fsnotify"
import (
	"syfiletoserver"
	"fmt"
)

func main() {
	//获取配置文件
	err,reCon:=syfiletoserver.Readconfig()
	if err!=nil{
		fmt.Println(err)
		return
	}
	//开始监控文件
	//-------------监控开始-----------------

	watch, _ := fsnotify.NewWatcher()
	w := syfiletoserver.Watch{
		watch,
	}
	w.WatchDir("E:/filetest",reCon);
	select {};
}