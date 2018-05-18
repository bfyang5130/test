package main

import "github.com/fsnotify/fsnotify"

func main() {
	//开始监控文件
	//-------------监控开始-----------------

	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}
	w.watchDir("E:/filetest", conn);
	select {};
}