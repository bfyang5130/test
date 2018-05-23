package syfiletoserver

import (
	"github.com/fsnotify/fsnotify"
	"time"
	"os"
	"fmt"
	"crypto/md5"
	"io"
	"path/filepath"
)

/**
 * 文件监控
 */
type Watch struct {
	Swatch *fsnotify.Watcher;
}

/**
 * 文件信息
 */
type sysFileInfo struct {
	fName  string
	fSize  int64
	fMtime time.Time
	fPerm  os.FileMode
	fMd5   string
	fType  bool
}

func Read1() []byte {
	path := "F:/test/flip.html"
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return chunks
}

//监控目录
func (w *Watch) WatchDir(dir string,reCon []SerConfig) {

	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//这里判断是否为目录，只需监控目录即可
		//目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir() {
			path, err := filepath.Abs(path);
			if err != nil {
				return err;
			}
			err = w.Swatch.Add(path);
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
			case ev := <-w.Swatch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println("创建文件 : ", ev.Name);
						//这里获取新创建文件的信息，如果是目录，则加入监控中
						fi, err := os.Stat(ev.Name);

						//conn.Write([]byte(newName))
						if err == nil && fi.IsDir() {
							w.Swatch.Add(ev.Name);
							fmt.Println("添加监控 : ", ev.Name);
							//读取文件信息
							//fInfo := getFileInfo(ev.Name)
							fmt.Printf("源站创建了新目录：%s\n", ev.Name)
							WriteOpLogFile("c",ev.Name,reCon)
						}else if err==nil{
							//读取文件信息
							//fInfo := getFileInfo(ev.Name)
							fmt.Printf("源站创建了新文件：%s\n", ev.Name)
							WriteOpLogFile("c",ev.Name,reCon)
						}

					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("写入文件 : ", ev.Name);
						//读取文件信息
						fmt.Printf("源站写入了文件：%s\n", ev.Name)
						WriteOpLogFile("w",ev.Name,reCon)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fmt.Println("删除文件 : ", ev.Name);
						//如果删除文件是目录，则移除监控
						fi, err := os.Stat(ev.Name);
						if err != nil {
							fmt.Printf("删除文件：%s\n", ev.Name)
							WriteOpLogFile("d",ev.Name,reCon)
						}
						if err == nil && fi.IsDir() {
							//文件删除后，不能再读取原来的文件，所以直接把文件名传送过去
							WriteOpLogFile("d",ev.Name,reCon)
							w.Swatch.Remove(ev.Name);
							fmt.Println("删除监控 : ", ev.Name);
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						//读取文件信息
						//fInfo:=getFileInfo( ev.Name)
						//newName :=fmt.Sprintf( "%s",fInfo.fName)
						//conn.Write([]byte(newName))
						//fmt.Println("重命名文件 : ", ev.Name);
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						w.Swatch.Remove(ev.Name);
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						//读取文件信息
						fInfo := getFileInfo(ev.Name)
						newName := fmt.Sprintf("%s", fInfo.fName)
						WriteOpLogFile("m",newName,reCon)
						fmt.Println("修改权限 : ", ev.Name);
					}
				}
			case err := <-w.Swatch.Errors:
				{
					fmt.Println("error : ", err);
					return;
				}
			}
		}
	}();
}

/**
 * 读取文件信息
 */
func getFileInfo(filename string) *sysFileInfo {
	fi, err := os.Lstat(filename)
	if err != nil {
		fmt.Println("info ERROR", err)
		return nil
	}
	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Println("open ERROR", err)
		return nil
	}
	defer fileHandle.Close()

	h := md5.New()
	_, err = io.Copy(h, fileHandle)
	fileInfo := &sysFileInfo{
		fName:  fi.Name(),
		fSize:  fi.Size(),
		fPerm:  fi.Mode().Perm(),
		fMtime: fi.ModTime(),
		fType:  fi.IsDir(),
		fMd5:   fmt.Sprintf("%x", h.Sum(nil)),
	}

	return fileInfo
}
