package main

import (
	"fmt"
	"github.com/google/gops/goprocess"
)

func main() {
	p:=goprocess.FindAll()
	if p==nil{
		fmt.Printf("没有进程\n")
		return
	}
	for _,v:=range p{

		fmt.Printf("%s",v.Exec)
	}
}
