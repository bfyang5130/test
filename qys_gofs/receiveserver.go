package main

import (
	"fmt"
	"net"
	"io"
)

func main(){
	fmt.Println("TCP服务器启动")
	//监听端口
	l,err :=net.Listen("tcp",":60010")
	if err!=nil{
		fmt.Printf("TCP端口创建出错:%s\n",err.Error())
	}
//接收连接信息并打印出控制台
   for {
	   c,err:=l.Accept()
   	 if err!=nil{
		 if ne,ok := err.( net.Error );ok && ne.Temporary(){
			 continue
		 }
		 fmt.Println( "network error",err )
	 }
	 go FitCon(c)
   }


}

func FitCon(conn net.Conn){
	defer conn.Close()

	for{
		buffer :=make( []byte,2048)
		num,err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println( "cannot read",err )
		}
		newString:=string(num)
		if newString!=""{
			fmt.Printf("%s",buffer[0:num])
		}

	}
}