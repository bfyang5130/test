package main

import (
	"net"
	"fmt"
	"os"
	"io"
	"time"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
)
/**
 * 文件监控
 */
type Watch struct {
	watch *fsnotify.Watcher;
}
/**
 * 文件信息
 */
type sysFileInfo struct{
	fName       string
	fSize       int64
	fMtime      time.Time
	fPerm       os.FileMode
	fMd5        string
	fType       bool
}

func main() {

	address := `127.0.0.1:60010`
	//开始连接服务器
	fmt.Printf("开始创建TCP连接,连接到：%s\n", address)
	conn, err := net.Dial("tcp", address)
	defer conn.Close()
	if err != nil {
		fmt.Printf("连接失败:%s\n", err.Error())
	}
	fmt.Println("连接成功")
	fmt.Println("开始传送数据...")
	//开始监控文件
	//-------------监控开始-----------------

	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}
	w.watchDir("E:/filetest",conn);
	select {};
	//-----------监控结束--------



	newFile:=Read1()
	//读取文件里的东西
	conn.Write(newFile)
}

func Read1() []byte{
	path:="F:/test/flip.html"
	fi,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte,1024,1024)
	buf := make([]byte,1024)
	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return chunks
}

//监控目录
func (w *Watch) watchDir(dir string,conn net.Conn) {

	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//这里判断是否为目录，只需监控目录即可
		//目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir() {
			path, err := filepath.Abs(path);
			if err != nil {
				return err;
			}
			err = w.watch.Add(path);
			if err != nil {
				return err;
			}
			fmt.Println("监控 : ", path);
		}
		return nil;
	});
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println("创建文件 : ", ev.Name);
						//这里获取新创建文件的信息，如果是目录，则加入监控中
						fi, err := os.Stat(ev.Name);
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name);
							fmt.Println("添加监控 : ", ev.Name);
						}
						conn.Write(string{ev.Name})
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("写入文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fmt.Println("删除文件 : ", ev.Name);
						//如果删除文件是目录，则移除监控
						fi, err := os.Stat(ev.Name);
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name);
							fmt.Println("删除监控 : ", ev.Name);
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						fmt.Println("重命名文件 : ", ev.Name);
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						w.watch.Remove(ev.Name);
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						fmt.Println("修改权限 : ", ev.Name);
					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err);
					return;
				}
			}
		}
	}();
}